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

var pattern *regexp.Regexp = regexp.MustCompile(`[Nn]once`)

func getNonce(ctx context.Context, address common.Address) (uint64, error) {
	if localNonce == nil {
		client := GetRpcClient()

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

func matchNonceError(errStr string) bool {
	res := pattern.FindStringSubmatch(errStr)
	return res != nil
}

func processSendingTxError(err error) error {
	if ok := matchNonceError(err.Error()); ok {
		localNonce = nil
	}
	return err
}
