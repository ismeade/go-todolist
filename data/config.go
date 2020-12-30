//
// config.go
// Copyright (C) 2020 liyang <yang.li@51vcheck.cn>
//
// Distributed under terms of the MIT license.
//

package data

var (
	bucketConfig []byte = []byte("config")
)

type Config struct {
}

func SetToken(token string) {
	insert(bucketConfig, "token", token)
}

func GetToken(token string) string {
	return read(bucketConfig, "token")
}

func SetKey(key string) {
	insert(bucketConfig, "key", key)
}

func GetKey(key string) string {
	return read(bucketConfig, "key")
}

func init() {
	initBucket(bucketConfig)
}
