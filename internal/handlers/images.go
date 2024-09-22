package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type ImageHandler struct {
}

func NewImageHandler() *ImageHandler {
	return &ImageHandler{}
}

// UploadHandler creates a new file in the images directory and copies the incoming file to this new file.
func (h *ImageHandler) UploadHandler(w http.ResponseWriter, r *http.Request) {
	const MAX_UPLOAD_SIZE = 10 << 20 // Set the max upload size to 10 MB
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(w, "The uploaded file is too big. Please choose a file that's less than 10MB in size", http.StatusBadRequest)
		return
	}
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	// Create the uploads folder if it doesn't already exist
	err = os.MkdirAll("./images", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Create a new file in the uploads directory
	filename := fmt.Sprintf("/images/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename))
	dst, err := os.Create("." + filename)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	// Copy the uploaded file to  the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response, _ := json.Marshal(map[string]string{"location": filename})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ServeImages serve image directory as static files.
func (h *ImageHandler) ServeImages(w http.ResponseWriter, r *http.Request) {
	fs := http.StripPrefix("/images/", http.FileServer(http.Dir("./images")))
	fs.ServeHTTP(w, r)
}
