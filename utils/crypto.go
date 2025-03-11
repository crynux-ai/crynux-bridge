package utils

import (
	"crypto/ecdsa"
	"errors"

	"github.com/ethereum/go-ethereum/crypto"
)

func GetPubKeyFromPrivKey(privKey string) (string, error) {
	privKeyECDSA, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return "", err
	}

	pubKeyCrypto := privKeyECDSA.Public()
	pubKeyECDSA, ok := pubKeyCrypto.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("error casting public key to ECDSA")
	}

	pubKeyBytes := crypto.FromECDSAPub(pubKeyECDSA)
	if len(pubKeyBytes) != 65 {
		return "", errors.New("umcompressed public key bytes length is not 65")
	}
	pubKeyBytes = pubKeyBytes[1:]

	return string(pubKeyBytes), nil
}
