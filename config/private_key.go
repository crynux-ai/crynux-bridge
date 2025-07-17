package config

import (
	"os"
	"strings"
)

func GetPrivateKey(file string) string {
	b, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
