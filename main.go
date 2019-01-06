package main

import (
	"net/http"
	logger "github.com/Sirupsen/logrus"
	"hlcup/domain"
)

//
//type router struct {
//
//}
//
type Server struct {
	log   *logger.Logger
	store map[int]domain.Account
}
//
//func NewRouter() *router {
//	return &router{
//		// empty
//	}
//}
//
func NewServer() *Server {
	return &Server{
		log:   logger.New(),
		store: make(map[int]domain.Account),
	}
}
//
//func (s *Server) serve() {
//
//}


func main() {
	var s = NewServer()

	var err = s.loadInitialData()
	if nil != err {
		s.log.Fatal("Could not load initial data")
		panic(err)
	}

	http.HandleFunc("/accounts/new", s.createAccount)
	http.HandleFunc("/accounts/likes", s.addLikes)

	s.log.Info("starting dating service")
	s.log.Fatal(http.ListenAndServe(":8080", nil))
}
