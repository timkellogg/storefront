package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/micro/go-micro/client"
	accountProto "github.com/timkellogg/store/account/protos/account"
)

// API - application struct that implements RPC calls
type API struct{}

// getAccountHandler - returns user account information
func (a *API) getAccountHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID := vars["accountID"]

	accountService := accountProto.AccountsServiceClient("go.micro.srv.accounts", client.DefaultClient)
	account, err := accountService.GetAccount(context.Background(), &accountProto.GetAccountRequest{Id: accountID})
	if err != nil {
		log.Println(err)
	}

	response, err := json.Marshal(account)
	if err != nil {
		log.Println(err)
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"account": string(response)})
}

// createAccountHandler - builds user account
func (a *API) createAccountHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	password := r.Form.Get("password")
	email := r.Form.Get("email")
	name := r.Form.Get("name")

	account := accountProto.Account{Password: password, Email: email, Name: name}

	accountService := accountProto.AccountsServiceClient("go.micro.srv.accounts", client.DefaultClient)
	_, err := accountService.CreateAccount(context.Background(), &account)
	if err != nil {
		log.Println(err)
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{})
}
