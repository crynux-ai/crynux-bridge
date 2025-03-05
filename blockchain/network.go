package blockchain

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func GetQueuedTasks(ctx context.Context) (*big.Int, error) {
	netstatsInstance := GetNetstatsContractInstance()

	callCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	opts := &bind.CallOpts{
		Pending: false,
		Context: callCtx,
	}

	return netstatsInstance.QueuedTasks(opts)
}
