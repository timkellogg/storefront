package main

import (
	"log"

	proto "github.com/timkellogg/store/user/protos/user"
	"gopkg.in/mgo.v2"
)

// UserRepository - collection of identities/users
type UserRepository struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	// COLLECTION - resource name
	COLLECTION = "identities"
)

// Connect - establish database connection
func (i *UserRepository) Connect() {
	session, err := mgo.Dial(i.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(i.Database)
}

// FindByID - returns a single user by id
func (i *UserRepository) FindByID(id string) (proto.User, error) {
	var user proto.User
	err := db.C(COLLECTION).FindId(id).One(&user)
	return user, err
}
