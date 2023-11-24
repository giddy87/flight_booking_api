package models

import (
	"github.com/giddy87/flight_booking_api/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Flight struct {
	gorm.Model
	//Flight_ID        uint `gorm:"primaryKey"`
	//CreatedAt        time.Time
	Name             string `json:"name"`
	Tail_number      string `json:"tail_number"`
	Flight_number    int64  `json:"flight_number"`
	Departure_city   string `json:"departure_city"`
	Destination_city string `json:"destination_city"`
	//Date                 *time.Time  `json:"time"`
	//Journey_time         *time.Timer `json:"journey_time"`
	Seats                int64   `json:"seats"`
	Seats_left           int64   `json:"seats_left"`
	Business_class_price float64 `json:"business_class_price"`
	Economy_class_price  float64 `json:"economy_class_price"`
}

type Passenger struct {
	gorm.Model
	//ID              uint   `gorm:"primaryKey"`
	Name            string `json:"first_name"`
	Middle_name     string `json:"middle_name"`
	Last_name       string `json:"last_name"`
	Passport_number string `json:"passport_number"`
	Age             int8   `json:"age"`
	User_id         int64
}

type Passenger_booking struct {
	gorm.Model
	Name             string `json:"first_name"`
	Middle_name      string `json:"middle_name"`
	Last_name        string `json:"last_name"`
	Passport_number  string `json:"passport_number"`
	User_id          int64  `json:"user_id"`
	Age              int8   `json:"age"`
	Booking_Id       int64
	Flight_Name      string  `json:"flight_name"`
	Tail_number      string  `json:"tail_number"`
	Flight_Number    int64   `"json:"flight_number"`
	Departure_city   string  `json:"departure_city"`
	Destination_city string  `json:"destination_city"`
	Seat_number      int64   `json:"seat_number"`
	Ticket_tier      int8    `json:"ticket_tier"` // 0 for economy, 1 for business
	Paid             bool    `json:"paid"`
	Price            float64 `json:"price"`
}

type Card struct {
	Name            string
	Billing_Address string
	Number          string
	CVV             int8
	Expiry          string
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Passenger{})
	db.AutoMigrate(&Flight{})
	db.AutoMigrate(&Passenger_booking{})
}

func (p *Passenger) CreatePassenger() *Passenger {
	db.NewRecord(p)
	db.Create(&p)
	return p
}
func (p *Flight) CreateFlight() *Flight {
	db.NewRecord(p)
	db.Create(&p)
	return p
}
func (p *Passenger_booking) CreateBooking() *Passenger_booking {
	db.NewRecord(p)
	db.Create(&p)
	return p
}

func GetAllPassengers() []Passenger {
	var Passengers []Passenger
	db.Find(&Passengers)
	return Passengers
}

func GetAllFlights() []Flight {
	var Flights []Flight
	db.Find(&Flights)
	return Flights
}
func GetAllBookings() []Passenger_booking {
	var Bookings []Passenger_booking
	db.Find(&Bookings)
	return Bookings
}
func GetAllPassengersOnAFlight(ID int64) []Passenger_booking {
	var Passengers_booked []Passenger_booking
	db.Where("Passengers_booked.flight_details.ID=?", ID).Find(&Passengers_booked)
	return Passengers_booked
}

func GetPassengerById(ID int64) (*Passenger, *gorm.DB) {
	var getPassenger Passenger
	db := db.Where("ID=?", ID).Find(&getPassenger)
	return &getPassenger, db
}
func GetPassengerByUserId(ID int64) (*Passenger, *gorm.DB) {
	var getPassenger Passenger
	db := db.Where("User_id=?", ID).Find(&getPassenger)
	return &getPassenger, db
}
func GetFlightById(ID int64) (*Flight, *gorm.DB) {
	var getFlight Flight
	db := db.Where("ID=?", ID).Find(&getFlight)
	return &getFlight, db
}
func GetFlightByFlightNumber(number int64) (*Flight, *gorm.DB) {
	var getFlight Flight
	db := db.Where("Flight_number=?", number).Find(&getFlight)
	return &getFlight, db
}

func DeletePassenger(Id int64) Passenger {
	var pass Passenger
	db.Where("ID = ?", Id).Delete(pass)
	return pass
}
func FindFlight(dep string, dest string) ([]Flight, *gorm.DB) {
	var getFlight []Flight
	db := db.Where("Departure_city = ?", dep).Where(db.Where("Destination_city = ?", dest)).Find(&getFlight)
	return getFlight, db
}

func GetBookingById(ID int64) (*Passenger_booking, *gorm.DB) {
	var Booking Passenger_booking
	db := db.Where("Booking_Id=?", ID).First(&Booking)
	return &Booking, db
}
