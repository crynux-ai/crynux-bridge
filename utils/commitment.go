package utils

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func HexStrToCommitment(s string) (*[32]byte, error) {
	bs, err := hexutil.Decode(s)
	if err != nil {
		return nil, err
	}
	if len(bs) != 32 {
		return nil, err
	}
	res := (*[32]byte)(bs)
	return res, nil
}
