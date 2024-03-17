package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	tempImagesFolder = "temp-images"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	//Max size 10MB bitwise "<<"
	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retriveing the file")
		fmt.Println(err)
		return
	}
	defer file.Close()
	if _, err := os.Stat(tempImagesFolder); os.IsNotExist(err) {
		os.Mkdir(tempImagesFolder, os.ModePerm)
	}

	tempFile, err := os.CreateTemp(tempImagesFolder, "upload-*.png")
	if err != nil {
		log.Fatalf("Error Creating file %s", err)
	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)
	resp, err := JsonResp("Successfully Uploaded")
	if err != nil {
		log.Fatalf("An Error Occurred: %s", err)
	}
	w.Write(resp)
}

func JsonResp(msg string, data ...map[string]interface{}) ([]byte, error) {
	resp := make(map[string]string)
	resp["message"] = "SuccessFully Uploaded file"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	return jsonResp, nil
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Returned from Running server %s \n", r.URL.Path)
	})
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Upload Service")
	setupRoutes()
}
