package main

import (
	"log"

	proto "github.com/timkellogg/store/account/protos/account"
	"gopkg.in/mgo.v2"
)

// AccountRepository - collection of accounts
type AccountRepository struct {
	Server   string
	Database string
}

var db *mgo.Database

const collection = "accounts"

// Connect - establish database connection
func (i *AccountRepository) Connect() {
	session, err := mgo.Dial(i.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(i.Database)
}

// FindByID - returns a single account by id
func (i *AccountRepository) FindByID(id string) (proto.Account, error) {
	var account proto.Account
	err := db.C(collection).FindId(id).One(&account)
	return account, err
}
