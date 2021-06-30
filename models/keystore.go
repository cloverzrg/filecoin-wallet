package models

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/jinzhu/gorm"
)

type KeyStore struct {
	gorm.Model
	Type       types.KeyType   `json:"type"`
	PrivateKey string          `json:"private_key"` // hex encode
	PublicKey  string          `json:"public_key"`
	Address    address.Address `json:"address"`
}
