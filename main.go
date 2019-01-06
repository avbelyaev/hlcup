package main

import (
	"net/http"
	logger "github.com/Sirupsen/logrus"
	"hlcup/domain"
	"github.com/gorilla/mux"
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
	r.HandleFunc("/accounts/likes", s.addLikes).Methods("POST")
	http.Handle("/", r)
}

func main() {
	var s = NewServer()
	s.setupRouter()

	var err = s.loadInitialData()
	if nil != err {
		s.log.Fatal("Could not load initial data. Terminating")
		panic(err)
	}

	s.log.Info("starting dating service")
	s.log.Fatal(http.ListenAndServe(":8080", nil))
}
