package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var currentTime = time.Now()

func main() {
	rand.Seed(time.Now().UnixNano())
	Years := [2]int{currentTime.Year(), currentTime.Year() + 1}
	fmt.Println(Years)
	months := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	currentMonth := int(currentTime.Month())
	Months_left := months[currentMonth-1 : 12]
	fmt.Println(Months_left)
	days := []int{}
	for i := 1; i <= 30; i++ {
		days = append(days, i)

	}
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

	fmt.Println(Date)
	fmt.Println(string_date)
	view := strings.Join(string_date, "")
	fmt.Println(view)
	flight_time, error := time.Parse("2006-01-02", view)
	if error == nil {
		fmt.Println(flight_time)
	}
	//fmt.Println(flight_time)
	//formatted2 := currentTime.Format("2006-01-02")
	//fmt.Println(formatted2)

}
