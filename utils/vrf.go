package utils

import "math/big"

func VrfNeedValidation(vrfNumber []byte) bool {
	number := big.NewInt(0).SetBytes(vrfNumber)
	r := big.NewInt(0).Mod(number, big.NewInt(10)).Uint64()
	return r == 0
}
