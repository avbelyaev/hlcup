package main

import (
	"net/http"
)

type router struct {

}

type Server struct {
	router *router
}

func NewRouter() *router {
	return &router{
		// empty
	}
}

func NewServer() *Server {
	return &Server{
		router: NewRouter(),
	}
}

func (s *Server) serve() {
	http.HandleFunc("/", s.handleFuck())
	http.ListenAndServe(":8080", nil)
}


func main() {
	println("start")

	var s = NewServer()
	s.serve()
}
