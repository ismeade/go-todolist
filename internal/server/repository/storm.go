package repository

import (
	"go-todolist/util"
	"log"
	"os"
	"path"

	"github.com/asdine/storm"
)

// ConnectStorm Create database connection
func connectStorm(dbName string) *storm.DB {
	if dbName == "" {
		dbName = "default.db"
	}
	dbPath := util.GetEnvStr("DB_FILE", "~/.go-todolist/"+dbName)

	createDirIfNotExist(path.Dir(dbPath))

	db, err := storm.Open(dbPath)
	if err != nil {
		log.Fatal("FATAL ERROR: Exiting program! - Could not connect Embedded Database File")
	}
	return db
}

// createDirIfNotExist creates a directory if not found
func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}
