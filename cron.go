package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

func main() {
	loc, err := time.LoadLocation("Africa/Lagos")
	if err != nil {
		panic(err)
	}
	fmt.Println(loc)
	//cronJob := cron.NewWithLocation(loc)
	cronJob := cron.New()
	cronJob.AddFunc("* * * * * *", func() {
		fmt.Println("you are cute")
	})
	cronJob.Start()
	select {}
}
