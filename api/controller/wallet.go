package controller

import (
	"github.com/cloverzrg/filecoin-wallet/db"
	"github.com/cloverzrg/filecoin-wallet/filecoin"
	"github.com/cloverzrg/filecoin-wallet/logger"
	"github.com/cloverzrg/filecoin-wallet/models"
	"github.com/filecoin-project/go-address"
	"github.com/gin-gonic/gin"
)

func AddressDetail(c *gin.Context) {
	addr := c.Param("address")

	data := models.KeyStore{}
	err := db.DB.Where("address = ?", addr).First(&data).Error
	if err != nil {
		c.JSON(500, err)
		return
	}
	fromString, err := address.NewFromString(data.Address)
	if err != nil {
		c.JSON(500, err)
		return
	}

	fil, err := filecoin.WalletBalanceCache(fromString)
	if err != nil {
		c.JSON(500, err)
		return
	}

	messages, err := filecoin.StateListMessagesCache(fromString)
	if err != nil {
		c.JSON(500, err)
		return
	}
	logger.Info("StateListMessages size:", len(messages))
	c.HTML(200, "address_detail.tmpl", gin.H{
		"data":     data,
		"balance":  fil.String(),
		"messages": messages,
	})
}
