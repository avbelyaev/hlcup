package main

import (
	"fmt"
	logger "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"hlcup/domain"
	"net/http"
)

type Server struct {
	log   *logger.Logger
	store map[int]domain.Account
}

func NewServer() *Server {
	return &Server{
		log:   logger.New(),
		store: make(map[int]domain.Account),
	}
}

func (s *Server) setupRouter() {
	var r = mux.NewRouter()
	r.HandleFunc("/accounts/new", s.createAccount).Methods("POST")
	r.HandleFunc("/accounts/{id:[0-9]+}", s.updateAccount).Methods("POST")
	r.HandleFunc("/accounts/{id:[0-9]+}", s.getAccount).Methods("GET")
	r.HandleFunc("/accounts/likes", s.addLikes).Methods("POST")
	r.HandleFunc("/accounts/filter", s.filterAccounts).Methods("GET")
	http.Handle("/", r)
}

func main() {
	var s = NewServer()
	s.setupRouter()

	var err = s.loadInitialData()
	if nil != err {
		s.log.Fatal(fmt.Sprintf("could not load initial data: %+v", err))
		panic(err)
	}
	s.log.Info(fmt.Sprintf("number of accounts loaded: %d", len(s.store)))

	s.log.Info("starting dating service")
	s.log.Fatal(http.ListenAndServe(":8080", nil))
}
