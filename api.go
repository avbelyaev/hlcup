package main

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"hlcup/domain"
	"net/http"
)

func (s *Server) createAccount(writer http.ResponseWriter, request *http.Request) {
	// decode
	var newAccount domain.Account
	var err = decode(request, &newAccount)
	if nil != err {
		badRequest(writer, err)
		return
	}

	// validate
	err = newAccount.Validate()
	if nil != err {
		badRequest(writer, err)
		return
	}

	// save
	s.store[newAccount.ID] = newAccount

	// rsp
	writer.WriteHeader(201)
}

func (s *Server) addLikes(writer http.ResponseWriter, request *http.Request) {
	// decode
	var newLikes domain.Likes
	var err = decode(request, &newLikes)
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
}

func badRequest(writer http.ResponseWriter, err error) {
	log.WithError(err).Error("bad rq")
	http.Error(writer, "go fck yourself", 400)
}

func decode(request *http.Request, container interface{}) error {
	var decoder = json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(container)
}