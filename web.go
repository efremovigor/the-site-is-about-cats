package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

type JsonResponse struct {
	Ok   bool   `json:"ok"`
	Name string `json:"name"`
}

type KittensCatalogJsonResponse struct {
	Kittens []*KittenView `json:"kittens"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	tmpl, _ := template.ParseFiles(templatePath + "index.html")
	tmpl.Execute(w, "")
}

func ApiTopicSender(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(JsonResponse{Ok: true, Name: r.PostFormValue("name")})

	file, fileHeaders, err := r.FormFile("fileupload")
	if err != nil {
		return
	}

	defer file.Close()

	// copy example
	f, err := os.OpenFile(storageTmpFilePath+fileHeaders.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	io.Copy(f, file)

	fmt.Fprintln(w, string(data))
}

func ApiGetKittens(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(KittensCatalogJsonResponse{Kittens: getKittensCatalog()})
	fmt.Fprintln(w, string(data))

}

func createWebServer() {
	router := mux.NewRouter()
	router.HandleFunc("/", IndexHandler).Methods(http.MethodGet)
	router.HandleFunc("/api/topic/send", ApiTopicSender).Methods(http.MethodPost)
	router.HandleFunc("/api/catalog", ApiGetKittens).Methods(http.MethodGet)

	var dir string
	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(publicPath+http.Dir(dir))))

	srv := &http.Server{
		ReadTimeout:  readTimeoutRequest,
		WriteTimeout: writeTimeoutRequest,
		Addr:         socket,
		Handler:      router,
	}
	log.Fatal(srv.ListenAndServe())
}
