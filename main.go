package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/cloverzrg/filecoin-wallet/api"
	"github.com/cloverzrg/filecoin-wallet/db"
	"github.com/cloverzrg/filecoin-wallet/filecoin"
	"github.com/cloverzrg/filecoin-wallet/logger"
	"github.com/cloverzrg/filecoin-wallet/models"
	"github.com/filecoin-project/go-address"
	"os"
)

var (
	BuildTime  string
	GoVersion  string
	GitMessage string
)

func main() {
	err := api.Start()
	if err != nil {
		logger.Panic(err)
	}
}

func init() {
	type cmdParams struct {
		IsPrintVersion bool
	}
	params := &cmdParams{}
	flag.BoolVar(&params.IsPrintVersion, "v", false, "print version")
	flag.Parse()
	if params.IsPrintVersion {
		os.Exit(0)
	}
	msg := fmt.Sprintf("BuildAt: %s\nBuildBy: %s\nGit：%s", BuildTime, GoVersion, GitMessage)
	fmt.Println(msg)
	err := db.Connect()
	if err != nil {
		logger.Panic(err)
	}
	err = db.DB.AutoMigrate(&models.KeyStore{}).Error
	if err != nil {
		logger.Panic(err)
	}

	networkName, err := filecoin.Client.StateNetworkName(context.Background())
	if err != nil {
		logger.Panic(err)
	}
	logger.Info("当前连接网络:", networkName)
	if networkName == "mainnet" {
		address.CurrentNetwork = address.Mainnet
	} else {
		address.CurrentNetwork = address.Testnet
	}
}
