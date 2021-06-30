package main

import (
	"github.com/cloverzrg/filecoin-wallet/api"
	"github.com/cloverzrg/filecoin-wallet/db"
)

func main() {
	api.Start()
}

func init() {
	db.Connect()
}