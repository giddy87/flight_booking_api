// package datagenerator
package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var Airlines = [7]string{"Dana", "Arik", "Max Air", "Azman Air", "Ibom Air", "Green Africa", "Value Jet"}
var Locations = [8]string{"LAG", "ENU", "KAN", "PHC", "ABJ", "BEN", "BRN", "JOS"}
var Seats = 150
var digits = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
var alpha = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N"}
var ID = 1

func GenerateFlightData() []string {
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
	return data

}
func main() {

	//start := time.Now()
	f, err := os.OpenFile("flights.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	init_line := []string{"ID", "Name", "Tail_number", "Flight_number", "Departure_city", "Destination_city", "Seats", "Seats_left", "Business_class_price", "Economy_class_price"}
	w := csv.NewWriter(f)
	w.Write(init_line)
	for i := 0; i < 1200; i++ {
		w.Write(GenerateFlightData())
	}
	w.Flush()
	//elapsed := time.Since(start)
	//log.Printf("Binomial took %s", elapsed)
}
