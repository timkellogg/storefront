package main

import (
	"log"

	identificationProto "github.com/timkellogg/store/identification/protos/identification"
	"gopkg.in/mgo.v2"
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
	log.Println("Establishing identification service database connection...")
	session, err := mgo.Dial(i.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(i.Database)
	log.Println("Identification service database connection successful")
}

// FindByID - returns a single identity by id
func (i *IdentityRepository) FindByID(id string) (*identificationProto.Identity, error) {
	log.Printf("IdentityRepository FindByID: %s", id)
	var identity identificationProto.Identity
	err := db.C(COLLECTION).FindId(id).One(&identity)
	return &identity, err
}

// Create - saves a new identity
func (i *IdentityRepository) Create(identity *identificationProto.Identity) (*identificationProto.Identity, error) {
	log.Printf("IdentityRepository CreateIdentity: %v", identity)
	err := db.C(COLLECTION).Insert(&identity)
	return identity, err
}
