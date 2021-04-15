// Package storm provides ...
package storm

import (
	"go-todolist/util"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/asdine/storm"
	"github.com/mitchellh/go-homedir"
)

// ConnectStorm Create database connection
func ConnectStorm() *storm.DB {
	dbPath := util.GetEnvStr("DB_FILE", "")
	var err error

	if dbPath == "" {
		// Try in home dir
		dbPath, err = homedir.Expand("~/.geek-life/default.db")

		// If home dir is not detected, try in system tmp dir
		if err != nil {
			f, _ := ioutil.TempFile("geek-life", "default.db")
			dbPath = f.Name()
		}
	}

	CreateDirIfNotExist(path.Dir(dbPath))

	db, openErr := storm.Open(dbPath)
	if openErr != nil {
		log.Fatal("FATAL ERROR: Exiting program! - Could not connect Embedded Database File")
	}
	return db
}

// CreateDirIfNotExist creates a directory if not found
func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}
