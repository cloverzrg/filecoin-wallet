package controller

import (
	"encoding/hex"
	"github.com/cloverzrg/filecoin-wallet/db"
	"github.com/cloverzrg/filecoin-wallet/logger"
	"github.com/cloverzrg/filecoin-wallet/models"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/wallet"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	list := []models.KeyStore{}
	err := db.DB.Find(&list).Error
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, list)
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
