package main

import (
	"net/http"
	"fmt"
)

func (s *Server) handleFuck() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(418)
		fmt.Fprintf(writer, "fucker!")
	}
}