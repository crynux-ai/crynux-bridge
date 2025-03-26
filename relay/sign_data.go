package relay

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func jsonRemarshal(bytes []byte) ([]byte, error) {
    var ifce interface{}
    err := json.Unmarshal(bytes, &ifce)
    if err != nil {
        return nil, err
    }
    return json.Marshal(ifce)
}

func SignData(data interface{}, privateKeyStr string) (timestamp int64, signature string, err error) {

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return 0, "", err
	}

	dataBytes, err = jsonRemarshal(dataBytes)
	if err != nil {
		return 0, "", err
	}

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return 0, "", err
	}

	timestamp = time.Now().Unix()
	timestampStr := strconv.FormatInt(timestamp, 10)
	timestampBytes := []byte(timestampStr)

	signBytes := append(dataBytes, timestampBytes...)

	log.Debugln("sign string: " + string(signBytes))

	dataHash := crypto.Keccak256Hash(signBytes)

	signatureBytes, err := crypto.Sign(dataHash.Bytes(), privateKey)
	if err != nil {
		return 0, "", err
	}

	signature = hexutil.Encode(signatureBytes)

	log.Debugln("signature: " + signature)

	return timestamp, signature, nil
}
