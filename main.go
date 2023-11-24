package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "kubernetes-workshop: ", log.LstdFlags)
	logger.Println("starting the kubernetes-workshop-application")
	logger.Println("username:", os.Getenv("user"))
	logger.Println("password:", os.Getenv("password"))

	mux := mux.NewRouter()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("status: ok\n"))
	}).Methods(http.MethodGet)
	http.ListenAndServe(":8080", mux)
}
