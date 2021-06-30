package filecoin

import (
	"context"
	"github.com/cloverzrg/filecoin-wallet/cache"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

func GetCurrentHeight() (int64, error) {
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

func StateListMessages(address2 address.Address) (messages []cid.Cid, err error) {
	if x, found := cache.StateListMessagesCache.Get(address2.String()); found {
		return x.([]cid.Cid), nil
	}
	currentHeight, err := GetCurrentHeight()
	if err != nil {
		return messages, err
	}
	height := currentHeight - 2880
	messages, err = Client.StateListMessages(context.Background(), &api.MessageMatch{To: address2, From: address2}, types.TipSetKey{}, abi.ChainEpoch(height))
	cache.StateListMessagesCache.Set(address2.String(), messages, cache.DefaultExpiration)
	return messages, err
}
