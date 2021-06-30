package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

const (
	// For use with functions that take an expiration time.
	// -1 则不过期
	NoExpiration time.Duration = -1
	// For use with functions that take an expiration time. Equivalent to
	// passing in the same expiration duration as was given to New() or
	// NewFrom() when the cache was created (e.g. 5 minutes.)
	// 使用set方式时，如果时间为0，则使用cache.New时设置的过期时间
	DefaultExpiration time.Duration = 0
)

var CommonCache *cache.Cache

var StateListMessagesCache *cache.Cache

var BalanceCache *cache.Cache

func init() {
	CommonCache = cache.New(30*time.Second, 60*time.Second)
	StateListMessagesCache = cache.New(30*time.Second, 120*time.Second)

	BalanceCache = cache.New(30*time.Second, 120*time.Second)
}
