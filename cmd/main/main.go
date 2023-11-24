package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/giddy87/flight_booking_api/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterAppRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Server starting on port 9000...")
	log.Fatal(http.ListenAndServe("0.0.0.0:9000", r))

}
