package blockchain

import (
	"crynux_bridge/config"

	"golang.org/x/time/rate"
)

var limiter *rate.Limiter

func getLimiter() *rate.Limiter {
	if limiter == nil {
		appConfig := config.GetConfig()
		limiter = rate.NewLimiter(rate.Limit(appConfig.Blockchain.RPS), int(appConfig.Blockchain.RPS))
	}
	return limiter
}
