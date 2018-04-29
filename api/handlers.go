package main

import (
	"context"
	"encoding/json"
	"net/http"
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

type API struct{}

// // Get - find user by id
// func (s *Server) Get(ctx context.Context, req *proto.GetUserRequest, res *proto.User) error {
// 	user, err := userRepository.FindByID(req.Id)
// 	res.Email = user.Email
// 	res.Guid = user.Guid
// 	res.Name = user.Name
// 	return err
// }

func (a *API) Call(ctx context.Context, req *api.Request) {

}
