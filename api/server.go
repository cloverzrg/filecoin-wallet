package api

import (
	"github.com/cloverzrg/filecoin-wallet/api/controller"
	"github.com/gin-gonic/gin"
)

func Start() error {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	SetRouter(r)
	return r.Run(":80")
}

func SetRouter(r *gin.Engine) {
	r.GET("/", controller.Index)
	r.GET("/new", controller.NewKey)
}
