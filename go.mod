module github.com/cloverzrg/filecoin-wallet

go 1.16

require (
	github.com/filecoin-project/go-address v0.0.5
	github.com/filecoin-project/go-jsonrpc v0.1.4-0.20210217175800-45ea43ac2bec
	github.com/filecoin-project/go-state-types v0.1.1-0.20210506134452-99b279731c48
	github.com/filecoin-project/lotus v1.10.0
	github.com/gin-gonic/gin v1.7.2
	github.com/jinzhu/gorm v1.9.16
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/sirupsen/logrus v1.7.0
)

replace github.com/filecoin-project/filecoin-ffi => ./extern/filecoin-ffi
