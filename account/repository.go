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
func (a *AccountRepository) Connect() {
	session, err := mgo.Dial(a.Server)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	db = session.DB(a.Database)
}

// FindByID - returns a single account by id
func (a *AccountRepository) FindByID(id string) (proto.Account, error) {
	var account proto.Account
	err := db.C(collection).FindId(id).One(&account)
	return account, err
}

// Insert - add a record to the database
func (a *AccountRepository) Insert(account *proto.Account) error {
	err := db.C(collection).Insert(&account)
	return err
}
