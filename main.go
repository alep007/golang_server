package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var LocalFolderPath = "./storage/"

type Document struct {
	ID   string
	Name string
	Size int64
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/documents", getDocuments).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000", router))
}

func getDocuments(w http.ResponseWriter, r *http.Request) { //returns list of documents in json format
	var docs []Document
	docs = getDocumentsFromPath(LocalFolderPath)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(docs)
}
func getDocumentsFromPath(folderPath string) []Document { //get all files in the given path
	var docs []Document

	var fileList = getFiles(folderPath)
	for _, value := range fileList {

		absolutePath := folderPath + value
		verifyPath(absolutePath)
		hashMd5, err := hashFileMd5(absolutePath)
		fileSize := getSize(absolutePath)

		if err == nil {
			docs = append(docs,
				Document{ID: hashMd5, Name: value, Size: fileSize})
		}
	}
	return docs
}

/**
	this is the utils section, where there are methods for file management
**/
func verifyPath(absolutePath string) {
	if _, err := os.Stat(absolutePath); os.IsNotExist(err) {
		fmt.Println(err)
	}
}
func getSize(path string) int64 {
	file, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}
	sizeFile := file.Size()
	return sizeFile
}
func getFiles(path string) []string {
	var files []string
	location, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range location {
		files = append(files, f.Name())
	}
	return files
}
func hashFileMd5(filePath string) (string, error) {
	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil

}
