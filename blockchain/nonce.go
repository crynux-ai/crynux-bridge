package blockchain

import (
	"context"
	"strconv"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/common"
)

var txMutex sync.Mutex

func getNonce(ctx context.Context, address common.Address) (uint64, error) {
	client := GetRpcClient()

	if err := getLimiter().Wait(ctx); err != nil {
		return 0, err
	}

	callCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	nonce, err := client.PendingNonceAt(callCtx, address)
	if err != nil {
		return 0, err
	}
	log.Debugln("Nonce from blockchain: " + strconv.FormatUint(nonce, 10))
	return nonce, err
}
