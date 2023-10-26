package config

import "os"

func GetPrivateKey(file string) string {
	b, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func GetTestPrivateKey() string {
	return ""
}