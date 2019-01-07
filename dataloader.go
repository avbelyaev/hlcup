package main

import (
	"archive/zip"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
	"hlcup/domain"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	INITIAL_ZIP  = "data/data.zip"
	INITIAL_JSON = "data/example.json"
)

func (s *Server) loadInitialData() error {
	s.log.Info("loading initial data")

	// unzip data into dir
	var dest = "/tmp/hlcup"
	var err = unzip(INITIAL_ZIP, dest)
	if nil != err {
		return errors.Wrap(err, "could not unzip "+INITIAL_ZIP)
	}

	// prepare consumer to save data into memory
	var dataConsumer = func(data []byte) error {
		// parse json
		var initialAccounts domain.Accounts
		var err = json.Unmarshal(data, &initialAccounts)
		if nil != err {
			return errors.Wrap(err, "unmarshalling error")
		}

		// put into in-memory storage
		for _, account := range initialAccounts.Accounts {
			s.store[account.ID] = account
		}
		return nil
	}

	// parse each unzipped file
	err = readData(dest, dataConsumer)
	if nil != err {
		return errors.Wrap(err, "cloud not read jsons from "+dest)
	}

	s.log.Info("initial data successfully loaded")
	return nil
}

func readData(path string, consumer func(data []byte) error) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return errors.New("could not read dir " + path)
	}

	for _, f := range files {
		var fileContent, err = ioutil.ReadFile(path + "/" + f.Name())
		if nil != err {
			return errors.New("could not read file " + f.Name())
		}

		err = consumer(fileContent)
		if nil != err {
			log.Error("could not consume file " + f.Name() + ". Skipping")
		}
	}
	return nil
}

func unzip(src, dest string) error {
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
