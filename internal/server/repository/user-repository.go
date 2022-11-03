package repository

import (
	"errors"
	"go-todolist/internal/server/entity"
	"log"

	"github.com/asdine/storm"
	"golang.org/x/crypto/bcrypt"
)

//type UserRepository interface {
//CreateUser(user entity.User) entity.User
//UpdateUser(user entity.User) entity.User
//FindByUsername(username string) entity.User
//}

//type userConnection struct {
//db *storm.DB
//}

var (
	db *storm.DB = connectStorm("user.db")
)

func CreateUser(user entity.User) entity.User {
	// TODO
	user.Password = hashAndSalt([]byte(user.Password))
	db.Save(user)
	return user
}

func VerifyCredential(username string, password string) error {
	var user entity.User
	err := db.One("username", username, &user)
	if err != nil {
		return err
	}
	password = hashAndSalt([]byte(password))
	if password != user.Password {
		return errors.New("password error")
	}
	return nil
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
