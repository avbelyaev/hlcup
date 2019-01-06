package main

import (
	"archive/zip"
	"encoding/json"
	"github.com/pkg/errors"
	"hlcup/domain"
	"io"
	"os"
	"path/filepath"
)

const (
	INITIAL_ZIP  = "resources/resources.zip"
	INITIAL_JSON = "resources/example.json"
)

func (s *Server) loadInitialData() error {
	s.log.Info("loading initial data")

	// prepare consumer for unzipped json
	var consumer = func(data *os.File) error {
		// parse json
		var initialAccounts domain.Accounts
		var jsonParser = json.NewDecoder(data)
		if err := jsonParser.Decode(&initialAccounts); nil != err {
			return errors.Wrap(err, "init resources parsing error")
		}

		// put into in-memory storage
		for _, account := range initialAccounts.Accounts {
			s.store[account.ID] = account
		}
		return nil
	}

	// unzip data and copying into memory
	var err = unzip(INITIAL_ZIP, "/tmp", consumer)
	if nil != err {
		return errors.Wrap(err, "could not unzip "+INITIAL_ZIP)
	}

	s.log.Info("initial data successfully loaded")
	return nil
}

func unzip(src, dest string, fileConsumer func(jsonData *os.File) error) error {
	var r, err = zip.OpenReader(src)
	if nil != err {
		return errors.Wrap(err, "error opening zip reader")
	}
	defer func() {
		if err := r.Close(); nil != err {
			panic(err)
		}
	}()

	_ = os.MkdirAll(dest, 0755)

	var extractAndWriteFile = func(f *zip.File) error {
		var readCloser, err = f.Open()
		if nil != err {
			return errors.Wrap(err, "error opening zip "+f.Name)
		}
		defer func() {
			if err := readCloser.Close(); nil != err {
				panic(err)
			}
		}()

		var path = filepath.Join(dest, f.Name)

		{
			_ = os.MkdirAll(filepath.Dir(path), f.Mode())
			// set up permissions to access file
			var f, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if nil != err {
				return errors.Wrap(err, "error opening file "+f.Name())
			}
			defer func() {
				if err := f.Close(); nil != err {
					panic(err)
				}
			}()

			_, err = io.Copy(f, readCloser)
			if nil != err {
				return errors.Wrap(err, "error while copying "+f.Name())
			}

			// apply consumer to fill in-memory repo
			err = fileConsumer(f)
			if nil != err {
				return errors.Wrap(err, "file "+f.Name()+" cannot be consumed")
			}
		}
		return nil
	}

	for _, f := range r.File {
		var err = extractAndWriteFile(f)
		if nil != err {
			return errors.Wrap(err, "could not extract and write "+f.Name)
		}
	}

	return nil
}
