package models

import (
	"fmt"
	"math/big"
	"strings"
)

type BigInt struct {
	big.Int
}

func (i BigInt) MarshalText() ([]byte, error) {
	s := i.String()
	return fmt.Appendf(nil, "\"%s\"", s), nil
}

func (i *BigInt) UnmarshalText(data []byte) error {
	s := string(data)
	s = strings.Trim(s, "\"")

	var z big.Int
	_, ok := z.SetString(s, 10)
	if !ok {
		return fmt.Errorf("not a valid big integer: %s", s)
	}
	i.Int = z
	return nil
}

func (i BigInt) MarshalJSON() ([]byte, error) {
	return i.MarshalText()
}

func (i *BigInt) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	return i.UnmarshalText(data)
}
