package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func HandleFormFile(writer http.ResponseWriter, request *http.Request) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	curFile := filepath.Base(ex)
	err = request.ParseMultipartForm(10 * 2048)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ok":     false,
			"module": curFile,
		}).Error(err.Error())
		return
	}
	f, h, err := request.FormFile("file")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ok":     false,
			"module": curFile,
		}).Error(err.Error())
		return
	}
	defer f.Close()
	logrus.WithFields(logrus.Fields{
		"ok":     true,
		"header": h,
		"file":   f,
	}).Info()

	ext := path.Ext(h.Filename)
	tmpFile, err := ioutil.TempFile("/tmp", "*"+ext)
	defer tmpFile.Close()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ok":    false,
			"error": err.Error(),
		}).Error()
		return
	}
	bytes, err := io.ReadAll(f)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ok":    false,
			"error": err.Error(),
		}).Error()
		return
	}
	_, err = tmpFile.Write(bytes)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ok":    false,
			"error": err.Error(),
		}).Error()
		return
	}
	logrus.WithFields(logrus.Fields{
		"ok": true,
	}).Info("File " + tmpFile.Name() + " successfully wrote")
}
