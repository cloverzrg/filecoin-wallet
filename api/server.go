package api

import (
	"github.com/cloverzrg/filecoin-wallet/api/controller"
	"github.com/gin-gonic/gin"
)

func Start() error {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.LoadHTMLGlob("./templates/*")
	SetRouter(r)
	return r.Run(":8080")
}

func SetRouter(r *gin.Engine) {
	r.GET("/", controller.Index)
	r.GET("/new", controller.NewKey)
	r.POST("/import", controller.ImportKey)

	r.GET("/address/:address", controller.AddressDetail)

	r.POST("/send", controller.Send)
}
