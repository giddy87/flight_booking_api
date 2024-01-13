package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/giddy87/flight_booking_api/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/reiver/go-telnet"
)

var ch = make(chan string)
var DB_HOST = os.Getenv("DB_HOST")
var DB_PORT = os.Getenv("DB_PORT")

func main() {
	r := mux.NewRouter()
	routes.RegisterAppRoutes(r)
	http.Handle("/", r)
	fmt.Println("Testing for DB connectivity...")
	for {
		time.Sleep(5 * time.Second)
		go FindDB()
		v := <-ch
		if v == "Connected" {
			close(ch)
			break
		} else {
			fmt.Println(v)
		}
	}
	fmt.Println("Server starting on port 9000...")
	log.Fatal(http.ListenAndServe("0.0.0.0:9000", r))

}

func FindDB() {
	telnetsClient, err := telnet.DialTo(DB_HOST + ":" + DB_PORT)
	if err != nil {
		fmt.Println(err)
		ch <- "Retrying"
	} else {
		fmt.Println("Connection Found")
		ch <- "Connected"
		defer telnetsClient.Close()

	}
}
