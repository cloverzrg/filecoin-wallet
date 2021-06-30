package filecoin

import (
	"context"
	"encoding/hex"
	"github.com/cloverzrg/filecoin-wallet/cache"
	"github.com/cloverzrg/filecoin-wallet/db"
	"github.com/cloverzrg/filecoin-wallet/models"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/wallet"
	"github.com/ipfs/go-cid"
)

func GetCurrentHeightCache() (int64, error) {
	key := "current_height"
	if x, found := cache.CommonCache.Get(key); found {
		return x.(int64), nil
	}
	chainHead, err := Client.ChainHead(context.Background())
	if err != nil {
		return 0, err
	}
	cache.CommonCache.Set(key, int64(chainHead.Height()), cache.DefaultExpiration)
	return int64(chainHead.Height()), nil
}

func StateListMessagesCache(address2 address.Address) (messages []cid.Cid, err error) {
	if x, found := cache.StateListMessagesCache.Get(address2.String()); found {
		return x.([]cid.Cid), nil
	}
	currentHeight, err := GetCurrentHeightCache()
	if err != nil {
		return messages, err
	}
	height := currentHeight - 2880
	messages, err = Client.StateListMessages(context.Background(), &api.MessageMatch{To: address2}, types.TipSetKey{}, abi.ChainEpoch(height))
	if err != nil {
		return messages, err
	}

	messages2, err := Client.StateListMessages(context.Background(), &api.MessageMatch{From: address2}, types.TipSetKey{}, abi.ChainEpoch(height))
	if err != nil {
		return messages, err
	}

	messages = append(messages, messages2...)

	cache.StateListMessagesCache.Set(address2.String(), messages, cache.DefaultExpiration)
	return messages, err
}

func WalletBalanceCache(address2 address.Address) (fil types.FIL, err error) {
	if x, found := cache.BalanceCache.Get(address2.String()); found {
		return x.(types.FIL), nil
	}
	balance, err := Client.WalletBalance(context.Background(), address2)
	if err != nil {
		return fil, err
	}

	fil, err = types.ParseFIL(balance.String() + "attofil")
	if err != nil {
		return fil, err
	}
	cache.BalanceCache.Set(address2.String(), fil, cache.DefaultExpiration)
	return fil, err
}

func Send(fromAddr, toAddr, val string) (cid cid.Cid, err error) {
	fil, err := types.ParseFIL(val + "fil")
	if err != nil {
		return cid, err
	}

	keyData := models.KeyStore{}
	err = db.DB.Where("address = ?", fromAddr).First(&keyData).Error
	if err != nil {
		return cid, err
	}
	fromAddress, err := address.NewFromString(keyData.Address)
	if err != nil {
		return cid, err
	}

	toAddress, err := address.NewFromString(toAddr)
	if err != nil {
		return cid, err
	}

	nonce, err := Client.MpoolGetNonce(context.Background(), fromAddress)
	if err != nil {
		return cid, err
	}

	message := types.Message{
		Version:    0,
		To:         toAddress,
		From:       fromAddress,
		Nonce:      nonce,
		Value:      abi.NewTokenAmount(fil.Int64()),
		GasLimit:   0,
		GasFeeCap:  abi.TokenAmount{},
		GasPremium: abi.TokenAmount{},
		Method:     0,
		Params:     nil,
	}

	messageWithGas, err := Client.GasEstimateMessageGas(context.Background(), &message, nil, types.TipSetKey{})
	if err != nil {
		return cid, err
	}

	localWallet, err := wallet.NewWallet(wallet.NewMemKeyStore())
	if err != nil {
		return cid, err
	}
	keyBytes, err := hex.DecodeString(keyData.PrivateKey)
	_, err = localWallet.WalletImport(context.Background(), &types.KeyInfo{
		Type:       types.KTSecp256k1,
		PrivateKey: keyBytes,
	})

	if err != nil {
		return cid, err
	}

	sign, err := localWallet.WalletSign(context.Background(), fromAddress, messageWithGas.Cid().Bytes(), api.MsgMeta{})
	if err != nil {
		return
	}

	cid, err = Client.MpoolPush(context.Background(), &types.SignedMessage{
		Message:   *messageWithGas,
		Signature: *sign,
	})
	return cid, err
}
