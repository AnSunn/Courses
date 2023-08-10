package main

import (
	"fmt"
	"github.com/AnSunn/3_Bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Server is listening")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
