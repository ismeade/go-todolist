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

func insert(bucket, key, value []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		err := b.Put(key, value)
		return err
	})
}

func rm(bucket, key []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		err := b.Delete(key)
		return err
	})
}

func read(bucket, key []byte) []byte {
	var val []byte
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		val = b.Get(key)
		return nil
	})
	return val
}

func fetchAll(bucket []byte) {
	var vals [][]byte
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		cur := b.Cursor()
		for k, v := cur.First(); k != nil; k, v = cur.Next() {
			vals = append(vals, v)
		}
		return nil
	})
}
