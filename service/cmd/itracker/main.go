package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/n3integration/itracker/internal/api"
)

const (
	defaultDir  = "../ui/dist/inventorytracker"
	defaultPort = 8020
)

var (
	dir  string
	port int
)

func init() {
	flag.StringVar(&dir, "dir", defaultDir, "the base directory of the web assets")
	flag.IntVar(&port, "port", defaultPort, "the web server port")
	flag.Parse()
}

func main() {
	r := mux.NewRouter()

	apiRouter := r.PathPrefix("/api/v1").Subrouter()
	apiRouter.Handle("/health", http.HandlerFunc(api.GetHealthHandler)).Methods(http.MethodGet)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: r,
	}

	log.Fatal(server.ListenAndServe())
}
