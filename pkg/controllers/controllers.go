package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/giddy87/flight_booking_api/pkg/models"
	"github.com/giddy87/flight_booking_api/pkg/utils"
	"github.com/gorilla/mux"
)

var digits = []rune("01234567890")

func Generate_RandId(n int) int64 {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = digits[rand.Intn(len(digits))]
	}
	result, _ := strconv.Atoi(string(b))
	return int64(result)

}
func GetAllPassengers(w http.ResponseWriter, r *http.Request) {
	AllPassengers := models.GetAllPassengers()
	result, _ := json.Marshal(AllPassengers)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
func GetAllBookings(w http.ResponseWriter, r *http.Request) {
	AllBookings := models.GetAllBookings()
	result, _ := json.Marshal(AllBookings)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
func GetBookingById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	BookId := vars["BookId"]
	ID, err := strconv.ParseInt(BookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookData, _ := models.GetBookingById(ID)
	result, _ := json.Marshal(bookData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
func GetPassengerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	passId := vars["PassId"]
	ID, err := strconv.ParseInt(passId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	passData, _ := models.GetPassengerById(ID)
	result, _ := json.Marshal(passData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetPassengerByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["UserId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	passData, _ := models.GetPassengerByUserId(ID)
	result, _ := json.Marshal(passData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
func DeletePassengerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	PassId := vars["PassId"]
	ID, err := strconv.ParseInt(PassId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	DelPassenger, err := models.DeletePassenger(ID)
	if err != nil {
		result, _ := json.Marshal(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
	} else {
		result, _ := json.Marshal(DelPassenger)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}

func GetAllFlights(w http.ResponseWriter, r *http.Request) {
	AllFlights := models.GetAllFlights()
	result, _ := json.Marshal(AllFlights)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
func GetFlightById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	passId := vars["PassId"]
	ID, err := strconv.ParseInt(passId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	passData, _ := models.GetFlightById(ID)
	result, _ := json.Marshal(passData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
func GetFlightByFlightNumber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	flightNum := vars["FlightNum"]
	number, err := strconv.ParseInt(flightNum, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	passData, _ := models.GetFlightByFlightNumber(number)
	result, _ := json.Marshal(passData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
func CreatePassenger(w http.ResponseWriter, r *http.Request) {
	createPass := &models.Passenger{}
	utils.ParseBody(r, createPass)
	User_id := Generate_RandId(6)
	createPass.User_id = User_id
	p := createPass.CreatePassenger() //CreatePassenger() is the Passenger models Method
	result, _ := json.Marshal(p)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Passenger_booking{}
	utils.ParseBody(r, createBook)
	fmt.Println(createBook)
	num := createBook.User_id
	fl := createBook.Flight_Number
	Pass, _ := models.GetPassengerByUserId(num)
	Flight, _ := models.GetFlightByFlightNumber(fl)
	if Flight.Seats_left == 0 {
		w.Header().Set("Content-Type", "application/json")
		result, _ := json.Marshal("Can't Book this flight; No seats available")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
	} else {
		if fl != 0 && num != 0 {
			createBook.Flight_Name = Flight.Name
			createBook.Tail_number = Flight.Tail_number
			createBook.Departure_city = Flight.Departure_city
			createBook.Destination_city = Flight.Destination_city
			createBook.Name = Pass.Name
			createBook.Last_name = Pass.Last_name
			createBook.Passport_number = Pass.Passport_number
			createBook.Age = Pass.Age
			createBook.Middle_name = Pass.Middle_name
		}
		if createBook.Ticket_tier == 1 {
			createBook.Price = Flight.Business_class_price
		} else if createBook.Ticket_tier == 0 {
			createBook.Price = Flight.Economy_class_price
		}
		booking_id := Generate_RandId(7)
		createBook.Booking_Id = booking_id
		p := createBook.CreateBooking() //CreatePassenger() is the Passenger models Method
		result, _ := json.Marshal(p)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}

}
func CreateFlight(w http.ResponseWriter, r *http.Request) {
	FlightM := &models.Flight{}
	utils.ParseBody(r, FlightM)
	fmt.Println(FlightM)
	p := FlightM.CreateFlight() //CreatePassenger() is the Passenger models Method
	result, _ := json.Marshal(p)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}
func FindFlightByCity(w http.ResponseWriter, r *http.Request) {
	FlightM := &models.Flight{}
	utils.ParseBody(r, FlightM)
	dep := FlightM.Departure_city
	dest := FlightM.Destination_city
	p, _ := models.FindFlight(dep, dest) //CreatePassenger() is the Passenger models Method
	result, _ := json.Marshal(p)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
func PayBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Booking_Id := vars["Booking_Id"]
	ID, err := strconv.ParseInt(Booking_Id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	CardDetails := &models.Card{}
	utils.ParseBody(r, CardDetails)
	if len(CardDetails.Number) != 11 || CardDetails.Name == "" {
		result, _ := json.Marshal("Incorrect card details")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
	} else {
		original_record, db := models.GetBookingById(ID)
		flight_record, db := models.GetFlightByFlightNumber(original_record.Flight_Number)
		original_record.Paid = true
		flight_record.Seats_left = flight_record.Seats_left - 1
		original_record.Seat_number = flight_record.Seats - flight_record.Seats_left
		db.Save(&original_record)
		db.Save(&flight_record)
		result, _ := json.Marshal("Payment successful")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}

}
func FindFlightByDate(w http.ResponseWriter, r *http.Request) {
	d := &models.Date{}
	utils.ParseBody(r, d)
	//date, error := time.Parse("2006-01-02", FlightM.Flight_date)
	//if error != nil {
	//	log.Panic(error)
	//}
	date, fault := time.Parse("2006-01-02", d.Flight_date)
	if fault != nil {
		fmt.Println("Error while parsing date :", fault)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fail, _ := json.Marshal("Incorrect date details")
		w.Write(fail)
	} else {
		p, _ := models.FindFlightByDate(date) //CreatePassenger() is the Passenger models Method
		result, err := json.Marshal(p)
		if err != nil {
			fmt.Println(err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			fail, _ := json.Marshal("No flights available for selected date")
			w.Write(fail)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(result)
		}

	}
}

func DeleteFlightById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Delete_Id := vars["Delete_id"]
	ID, err := strconv.ParseInt(Delete_Id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	DelFlight, err := models.DeleteFlight(ID)
	if err != nil {
		result, _ := json.Marshal(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
	} else {
		result, _ := json.Marshal(DelFlight)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}

}
