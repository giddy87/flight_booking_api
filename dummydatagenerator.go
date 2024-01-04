// package datagenerator
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

//var ch = make(chan []string, 100)

var Airlines = [7]string{"Dana", "Arik", "Max Air", "Azman Air", "Ibom Air", "Green Africa", "Value Jet"}
var Locations = [8]string{"LAG", "ENU", "KAN", "PHC", "ABJ", "BEN", "BRN", "JOS"}
var Seats = 150
var digits = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
var alpha = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N"}
var ID = 1
var currentTime = time.Now()
var Years = [2]int{currentTime.Year(), currentTime.Year() + 1}
var months = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
var currentMonth = int(currentTime.Month())
var Months_left = months[currentMonth-1 : 12]
var days = []int{}

func GenerateFlightData() []string {
	for i := 1; i <= 30; i++ {
		days = append(days, i)

	}
	data := make([]string, 0)
	rand.Seed(time.Now().UnixNano())
	data = append(data, strconv.Itoa(ID))
	ID++
	A := Airlines[rand.Intn(len(Airlines))]
	data = append(data, A)

	//Generate Tail Number
	B := make([]string, 6)
	for i := range B {
		if i <= 1 {
			B[i] = digits[rand.Intn(len(digits))]
		} else if i == 2 {
			B[i] = "-"
		} else {
			B[i] = alpha[rand.Intn(len(alpha))]
		}

	}
	b := strings.Join(B, "")
	data = append(data, b)

	//Generate Flight Number
	C := make([]string, 5)
	for i := range C {
		C[i] = digits[rand.Intn(len(digits))]
	}
	c := strings.Join(C, "")
	data = append(data, c)

	//Generate Departure city
	D := Locations[rand.Intn(len(Locations))]
	data = append(data, D)

	//Generate Arrival City
	v := 0
	for v < len(Locations) {
		E := Locations[rand.Intn(len(Locations))]
		if E != D {
			data = append(data, E)
			break
		}
		v++
	}

	//Generate Seats randomize between 100 to 140
	F := (rand.Intn(50) + 100)
	data = append(data, strconv.Itoa(F))

	//Generate seats left
	G := strconv.Itoa(rand.Intn(F))
	data = append(data, G)

	//Generate Business Class Price
	min := 225000
	max := 380000
	h := (rand.Intn(max-min) + min)
	H := strconv.Itoa(h)
	data = append(data, H)
	//Generate Economy Class Price
	I := strconv.Itoa(int(float32(h) * 0.4))
	data = append(data, I)

	//Generate Flight Dates
	Date := make([]int, 3)
	for i := range Date {
		if i == 0 {
			Date[i] = Years[rand.Intn(len(Years))]
			continue
		}
		if i == 1 {
			if Date[0] == int(currentTime.Year()) {
				Date[i] = Months_left[rand.Intn(len(Months_left))]
				continue
			}
			if Date[0] == int(currentTime.Year()+1) {
				Date[i] = months[rand.Intn(len(months))]
				continue
			}

		}

		if i == 2 {
			Date[i] = days[rand.Intn(len(days))]
		}

	}
	string_date := make([]string, 5)
	for i := range string_date {
		if i == 0 {
			//Year
			string_date[i] = strconv.Itoa(Date[0])
		}
		if i == 1 {
			string_date[i] = "-"
		}
		if i == 2 {
			//Month
			if Date[1] <= 9 {
				string_date[i] = "0" + strconv.Itoa(Date[1])
			} else {
				string_date[i] = strconv.Itoa(Date[1])
			}
		}
		if i == 3 {
			string_date[i] = "-"
		}
		if i == 4 {
			//Day
			if Date[2] <= 9 {
				string_date[i] = "0" + strconv.Itoa(Date[2])
			} else {
				string_date[i] = strconv.Itoa(Date[2])
			}
		}
	}
	view := strings.Join(string_date, "")
	J, error := time.Parse("2006-01-02", view)
	if error == nil {
		data = append(data, J.String())
		//ch <- data
		return data
	} else {
		//fmt.Println(error)
		return nil
	}
	//ch <- data
	//return data
}
func main() {
	//defer close(ch)
	start := time.Now()
	f, err := os.OpenFile("flights.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	init_line := []string{"ID", "Name", "Tail_number", "Flight_number", "Departure_city", "Destination_city", "Seats", "Seats_left", "Business_class_price", "Economy_class_price", "Flight_date"}
	w := csv.NewWriter(f)
	w.Write(init_line)
	for i := 0; i < 1500; i++ {
		dat := GenerateFlightData()
		if dat == nil {
			continue
		}
		w.Write(dat)
		//w.Write(GenerateFlightData())
		//GenerateFlightData()
		//_, ok := <-ch
		//if ok == true {
		//w.Write(<-ch)
		//}

	}
	w.Flush()
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
