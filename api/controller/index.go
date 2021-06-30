package controller

import (
	"context"
	"encoding/hex"
	"github.com/cloverzrg/filecoin-wallet/config"
	"github.com/cloverzrg/filecoin-wallet/db"
	"github.com/cloverzrg/filecoin-wallet/filecoin"
	"github.com/cloverzrg/filecoin-wallet/logger"
	"github.com/cloverzrg/filecoin-wallet/models"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/wallet"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	list := []models.KeyStore{}
	err := db.DB.Order("id asc").Find(&list).Error
	if err != nil {
		c.JSON(500, err)
		return
	}
	version, err := filecoin.Client.Version(context.Background())
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.HTML(200, "index.tmpl", gin.H{
		"version":  version,
		"network":  address.CurrentNetwork,
		"endpoint": config.Endpoint,
		"list":     list,
	})
}

// bls secp256k1
func NewKey(c *gin.Context) {
	keyType := c.DefaultQuery("type", "secp256k1")
	t := types.KeyType(keyType)
	generateKey, err := wallet.GenerateKey(t)
	if err != nil {
		c.JSON(500, err)
		return
	}

	keyStore := models.KeyStore{
		Type:       generateKey.Type,
		PrivateKey: hex.EncodeToString(generateKey.PrivateKey),
		PublicKey:  hex.EncodeToString(generateKey.PublicKey),
		Address:    generateKey.Address.String(),
	}
	err = db.DB.Create(&keyStore).Error
	if err != nil {
		c.JSON(500, err)
		return
	}
	logger.Info("new address:", keyStore.Address)
	c.Redirect(307, "/")
}

func ImportKey(c *gin.Context) {
	keyType := c.PostForm("type")
	privateKey := c.PostForm("privateKey")
	t := types.KeyType(keyType)

	privateBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		logger.Error(err)
		c.String(500, err.Error())
		return
	}

	keyInfo := types.KeyInfo{
		Type:       t,
		PrivateKey: privateBytes,
	}
	generateKey, err := wallet.NewKey(keyInfo)
	if err != nil {
		c.JSON(500, err)
		return
	}

	keyStore := models.KeyStore{
		Type:       generateKey.Type,
		PrivateKey: hex.EncodeToString(generateKey.PrivateKey),
		PublicKey:  hex.EncodeToString(generateKey.PublicKey),
		Address:    generateKey.Address.String(),
	}
	err = db.DB.Create(&keyStore).Error
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.Redirect(307, "/")
}
