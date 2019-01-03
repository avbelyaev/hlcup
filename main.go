package main

import (
	"net/http"
	log "github.com/Sirupsen/logrus"
)

//
//type router struct {
//
//}
//
//type Server struct {
//	router *router
//}
//
//func NewRouter() *router {
//	return &router{
//		// empty
//	}
//}
//
//func NewServer() *Server {
//	return &Server{
//		router: NewRouter(),
//	}
//}
//
//func (s *Server) serve() {
//
//}


func main() {
	log.Info("starting")

	http.HandleFunc("/fuck", HandleFuck())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
