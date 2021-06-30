package filecoin

import (
	"context"
	"github.com/cloverzrg/filecoin-wallet/cache"
)

func GetCurrentHeight() (int64, error) {
	if x, found := cache.CurrentHeightCache.Get("height"); found {
		return x.(int64), nil
	}
	chainHead, err := Client.ChainHead(context.Background())
	if err != nil {
		return 0, err
	}
	cache.CurrentHeightCache.Set("height", int64(chainHead.Height()), cache.DefaultExpiration)
	return int64(chainHead.Height()), nil
}
