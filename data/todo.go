//
// todo.go
// Copyright (C) 2020 liyang <yang.li@51vcheck.cn>
//
// Distributed under terms of the MIT license.
//

package data

import (
	"fmt"
	"time"
)

var (
	bucket []byte = []byte("data")
)

type todo struct {
	topic   string
	desc    string
	created string
}

func Add(topic, desc string) {
	t := todo{topic, desc, time.Now().Format("2006-01-02 03:04:05")}
	insert(bucket, topic, t.Show())
	fmt.Printf("add:%s\n", t.Show())
}

func GetAll() {

}

func (t todo) Show() string {
	return t.created + ": " + t.topic + " - " + t.desc + "\n"
}

func init() {
	initBucket(bucket)
}
