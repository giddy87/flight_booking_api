package routes

import (
	"github.com/giddy87/flight_booking_api/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterAppRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/passenger", controllers.CreatePassenger).Methods("POST")
	router.HandleFunc("/api/passenger", controllers.GetAllPassengers).Methods("GET")
	router.HandleFunc("/api/passenger/{PassId}", controllers.GetPassengerById).Methods("GET")
	router.HandleFunc("/api/passenger/user/{userId}", controllers.GetPassengerByUserId).Methods("GET")
	router.HandleFunc("/api/flight/id/{PassId}", controllers.GetFlightById).Methods("GET")
	router.HandleFunc("/api/book", controllers.CreateBooking).Methods("POST")
	router.HandleFunc("/api/book", controllers.GetAllBookings).Methods("GET")
	router.HandleFunc("/api/book/{BookId}", controllers.GetBookingById).Methods("GET")
	router.HandleFunc("/api/book/pay/{Booking_Id}", controllers.PayBooking).Methods("POST") //Pay for Ticket Booking using dummy card details
	router.HandleFunc("/api/flight", controllers.CreateFlight).Methods("POST")
	router.HandleFunc("/api/flight/{FlightNum}", controllers.GetFlightByFlightNumber).Methods("GET")
	router.HandleFunc("/api/flight/find", controllers.FindFlight).Methods("POST")
	router.HandleFunc("/api/flight", controllers.GetAllFlights).Methods("GET")
	router.HandleFunc("/api/flight/date", controllers.FindFlightByDate).Methods("POST")

	//router.HandleFunc("/passenger/{PassId}", controllers.DeletePassenger).Methods("DELETE")
	//router.HandleFunc("/passenger/{PassId}", controllers.UpdatePassenger).Methods("PUT")

}
