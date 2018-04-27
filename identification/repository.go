package main

import (
	"log"

	identificationProto "github.com/timkellogg/store/identification/protos/identification"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// IdentityRepository - collection of identities/users
type IdentityRepository struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	// COLLECTION - resource name
	COLLECTION = "identities"
)

// Connect - establish database connection
func (i *IdentityRepository) Connect() {
	session, err := mgo.Dial(i.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(i.Database)
}

// FindByID - returns a single identity by id
func (i *IdentityRepository) FindByID(id string) (identificationProto.Identity, error) {
	var identity identificationProto.Identity
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&identity)
	return identity, err
}
