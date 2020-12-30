package data

import (
	"fmt"
	"go-todolist/util"
	"os"

	"github.com/boltdb/bolt"
)

var (
	db *bolt.DB
)

func init() {
	//创建bolt数据库本地文件
	home, _ := util.Home()
	fmt.Println(home)
	os.MkdirAll(home+"/.todo", os.ModePerm)
	dbc, err := bolt.Open(home+"/.todo/data.db", 0600, nil)
	if err != nil {
		fmt.Println("open err:", err)
		return
	} else {
		db = dbc
	}
}

func initBucket(bucket []byte) {
	//创建bucket
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		if b == nil {
			_, err := tx.CreateBucket(bucket)
			return err
		}
		return nil
	})
}

func insert(bucket []byte, key, value string) error {

	k := []byte(key)
	v := []byte(value)
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)

		err := b.Put(k, v)
		return err
	})
	return err
}

func rm(bucket []byte, key string) {
	k := []byte(key)
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		err := b.Delete(k)
		return err
	})
}

func read(bucket []byte, key string) string {
	k := []byte(key)
	var val []byte
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		val = b.Get(k)
		return nil
	})
	return string(val)
}

func fetchAll(bucket []byte) {
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		cur := b.Cursor()
		for k, v := cur.First(); k != nil; k, v = cur.Next() {
			fmt.Printf("key is %s,value is %s\n", k, v)
		}
		return nil
	})
}
