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

func SetToken(token []byte) {
	insert(bucketConfig, []byte("token"), token)
}

func GetToken(token []byte) []byte {
	return read(bucketConfig, []byte("token"))
}

func SetKey(key []byte) {
	insert(bucketConfig, []byte("key"), key)
}

func GetKey(key []byte) []byte {
	return read(bucketConfig, []byte("key"))
}

func init() {
	initBucket(bucketConfig)
}
