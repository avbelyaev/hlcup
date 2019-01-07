package main

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"hlcup/domain"
	"net/http"
	"strconv"
)

func (s *Server) filterAccounts(writer http.ResponseWriter, request *http.Request) {
	var account = s.store[10003]

	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(account)
}

func (s *Server) getAccount(writer http.ResponseWriter, request *http.Request) {
	var pathVars = mux.Vars(request)
	var idStr = pathVars["id"]
	accountId, _ := strconv.Atoi(idStr)

	// check if account is present
	var account, ok = s.store[accountId]
	if !ok {
		notFound(writer)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(account)
}

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
	s.store[*newAccount.ID] = newAccount

	// rsp
	writer.WriteHeader(201)
}

func (s *Server) updateAccount(writer http.ResponseWriter, request *http.Request) {
	var pathVars = mux.Vars(request)
	var idStr = pathVars["id"]
	accountId, _ := strconv.Atoi(idStr)

	// check if updatable account is present
	var account, ok = s.store[accountId]
	if !ok {
		notFound(writer)
		return
	}

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

	// finally update
	err = account.Update(&newAccount)
	if nil != err {
		badRequest(writer, err)
		return
	}

	// rsp
	writer.WriteHeader(202)
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

func notFound(writer http.ResponseWriter) {
	http.Error(writer, "", 404)
}

func decode(request *http.Request, container interface{}) error {
	var decoder = json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(container)
}