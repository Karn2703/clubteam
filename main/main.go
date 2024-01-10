package main

import (
	"log"
	"net/http"

	"clubteam/pkg/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.MembersRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":7500", r))
}
