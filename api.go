package main

import (
	"net/http"
	"hlcup/domain"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
)

func (s *Server) createAccount(writer http.ResponseWriter, request *http.Request) {
	// check method
	if http.MethodPost != request.Method {
		badRequest(writer, nil)
		return
	}

	var newAccount domain.Account
	var decoder = json.NewDecoder(request.Body)

	// decode
	var err = decoder.Decode(&newAccount)
	if nil != err {
		badRequest(writer, err)
		return
	}

	// save
	s.store[newAccount.ID] = newAccount

	// rsp
	writer.WriteHeader(201)
	writer.Write(nil)
}

func (s *Server) addLikes(writer http.ResponseWriter, request *http.Request) {
	// check method
	if http.MethodPost != request.Method {
		badRequest(writer, nil)
		return
	}

	var newLikes domain.Likes
	var decoder = json.NewDecoder(request.Body)

	// decode
	var err = decoder.Decode(&newLikes)
	if nil != err {
		badRequest(writer, err)
		return
	}

	// validate
	err = newLikes.Validate()
	if nil != err {
		badRequest(writer, err)
		return
	}
	// TODO validate that liker and likee exist

	// rsp
	writer.WriteHeader(202)
	writer.Write(nil)
}

func badRequest(writer http.ResponseWriter, err error) {
	log.WithError(err).Error("bad rq")
	http.Error(writer, "go fck yourself", 400)
	writer.Write(nil)
}
