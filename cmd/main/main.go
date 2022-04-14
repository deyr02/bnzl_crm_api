package main

import (
	"log"
	"net/http"

	"github.com/deyr02/bnzl_crm/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.Actvity_Category_Routes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8090", r))

}
