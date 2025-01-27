package blockchain

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/common"
)

var localNonce *uint64
var txMutex sync.Mutex

var pattern *regexp.Regexp = regexp.MustCompile(`invalid nonce; got (\d+), expected (\d+)`)

func getNonce(ctx context.Context, address common.Address) (uint64, error) {
	if localNonce == nil {
		client, err := GetRpcClient()
		if err != nil {
			return 0, err
		}

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
		localNonce = &nonce
	}
	return *localNonce, nil
}

func addNonce(nonce uint64) {
	if *localNonce != nonce {
		log.Panic(fmt.Sprintf("local nonce changed, local nonce: %d, nonce: %d", *localNonce, nonce))
	}
	(*localNonce)++
}

func matchNonceError(errStr string) (uint64, bool) {
	res := pattern.FindStringSubmatch(errStr)
	if res == nil {
		return 0, false
	}
	nonceStr := res[len(res)-1]
	if len(nonceStr) == 0 {
		return 0, false
	}
	nonce, _ := strconv.ParseUint(nonceStr, 10, 64)
	return nonce, true
}

func processSendingTxError(err error) error {
	if nonce, ok := matchNonceError(err.Error()); ok {
		*localNonce = nonce
	}
	return err
}
