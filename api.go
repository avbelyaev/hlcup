package main

import (
	"net/http"
	"hlcup/domain"
	"encoding/json"
)

func (s *Server) createAccount(writer http.ResponseWriter, request *http.Request) {
	var newAccount domain.Account
	var decoder = json.NewDecoder(request.Body)

	// decode
	var err = decoder.Decode(&newAccount)
	if nil != err {
		s.badRequest(writer, err)
		return
	}

	// save
	s.store[newAccount.ID] = newAccount

	// rsp
	writer.WriteHeader(200)
	writer.Write(nil)
}

func (s *Server) addLikes(writer http.ResponseWriter, request *http.Request) {
	var newLikes domain.Like // TODO likeS
	var decoder = json.NewDecoder(request.Body)

	// decode
	var err = decoder.Decode(&newLikes)
	if nil != err {
		s.badRequest(writer, err)
		return
	}

	// validate
	err = newLikes.Validate()
	if nil != err {
		s.badRequest(writer, err)
		return
	}

	// rsp
	writer.WriteHeader(200)
	writer.Write([]byte("good"))
}

func (s *Server) badRequest(writer http.ResponseWriter, err error) {
	s.log.WithError(err).Error("bad rq")
	http.Error(writer, "go fck yourself", 400)
	writer.Write(nil)
}
