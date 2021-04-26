package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func (api API) HandlerAppImageUpload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("img")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	// forward image to file_server to store and get url back
	// add reference to app in app_service
}
