package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	identificationProto "github.com/timkellogg/store/identification/protos/identification"
	"google.golang.org/grpc"
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

func getIdentityHandler(w http.ResponseWriter, r *http.Request) {
	// pass this as context to route handler
	conn, err := grpc.Dial(config.Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	identificationService := identificationProto.NewIdentificationServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	log.Println("Gateway registered identification service client")

	is, err := identificationService.Get(ctx, &identificationProto.GetRequest{Id: "1"})

	respondWithJSON(w, 200, map[string]string{"user": is.GetEmail()})
}

func createIdentityHandler(w http.ResponseWriter, r *http.Request) {

}
