package controller

import (
	"context"
	"github.com/cloverzrg/filecoin-wallet/db"
	"github.com/cloverzrg/filecoin-wallet/filecoin"
	"github.com/cloverzrg/filecoin-wallet/models"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
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
	balance, err := filecoin.Client.WalletBalance(context.Background(), fromString)
	if err != nil {
		c.JSON(500, err)
		return
	}
	// 最近一周，每周有20160个块产生
	messages, err := filecoin.Client.StateListMessages(context.Background(), &api.MessageMatch{To: fromString}, types.TipSetKey{}, 12000)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.HTML(200, "address_detail.tmpl", gin.H{
		"data":     data,
		"balance":  balance.String(),
		"messages": messages,
	})
}
