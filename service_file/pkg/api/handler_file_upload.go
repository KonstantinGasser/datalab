package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

func (api API) HandlerFileUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	file, handler, err := r.FormFile("file")
	if err != nil {
		logrus.Errorf("could not get file from request: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	prefix := fmt.Sprint(time.Now().Unix())
	// TODO: what if filename: file.name.png -> edge case
	rawname := strings.Split(handler.Filename, ".")
	filename := rawname[0] + "_" + prefix + "." + rawname[1]
	imgURL := "/resources/images/app/" + filename
	f, err := os.OpenFile("."+imgURL, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		logrus.Errorf("could not create new file: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	if _, err := io.Copy(f, file); err != nil {
		logrus.Errorf("could not write to new file: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"image_ref": "http://localhost:8000/images/app/" + filename,
	})
}
