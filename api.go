package main

import (
	"net/http"
"github.com/gorilla/mux"
)

type APIServer struct {
	ListenAndServeAdrr string
}

func NewAPIServer(listenandserveaddr string) *APIServer {
	return &APIServer{
		ListenAndServeAdrr: listenandserveaddr,
	}
}

func (s *APIServer) Run() {
	r:=mux.NewRouter()
	r.HandleFunc("/account",nil)


}

func (s *APIServer) HandleAcc(w http.ResponseWriter, r *http.Request) error {

	return nil
}
func (s *APIServer) HandleGetAcc(w http.ResponseWriter, r *http.Request) error {

	return nil
}
func (s *APIServer) HandleCreateAcc(w http.ResponseWriter, r *http.Request) error {

	return nil
}
func (s *APIServer) HandleDeleteAcc(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func (s *APIServer) HandleTransferMoney(w http.ResponseWriter, r *http.Request) error {

	return nil
}
