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

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// API - application struct that implements RPC calls
type API struct{}

func getAccountHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID := vars["accountID"]

	accountService := accountProto.AccountsServiceClient("go.micro.account.client", client.DefaultClient)
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
