package controllers

import (
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path"

	"github.com/ksubedi/go-web-seed/core"
)

//Static serves static files and if not found throws 404
type Static struct {
	Dir    string
	Router *core.Router
}

func (s *Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Read file if possible
	filePath := path.Join("./", s.Dir, r.URL.Path)

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println(err)
		ThrowNotFoundError(s, w, r)
	} else {
		// If it is a directory, try to open index.html
		if fileInfo.IsDir() {
			filePath = path.Join(filePath, "index.html")
		}

		fileData, err := ioutil.ReadFile(filePath)

		//If its still null, throw 404, if the file is read, read file
		if err != nil {
			fmt.Println(err)
			ThrowNotFoundError(s, w, r)
		} else {

			//Find mimetype and set the header to that
			// mimeType := http.DetectContentType(fileData)
			ext := path.Ext(filePath)
			mimeType := mime.TypeByExtension(ext)

			w.Header().Set("Content-Type", mimeType)
			fmt.Fprint(w, string(fileData))
		}

	}

}

func ThrowNotFoundError(s *Static, w http.ResponseWriter, r *http.Request) {
	//Use gorilla's notfound handler if possible
	if s.Router.NotFoundHandler != nil {
		s.Router.NotFoundHandler.ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
	}

}
