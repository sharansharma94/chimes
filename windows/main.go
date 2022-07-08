package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gopkg.in/toast.v1"
)

func chime(title string) {

	if len(title) <= 0 {
		title = "Take deep breath and move"
	}

	notification := toast.Notification{
		AppID:   "Chime",
		Title:   title,
		Message: "Take a break",
		Actions: []toast.Action{
			{"protocol", "Doing it", ""},
			{"protocol", "Not doing it", ""},
		},
	}

	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}

func addInterval(period string, d1 time.Time, dt time.Time, interval int64) time.Time {
	switch period {
	case "h":
		d1 = dt.Add(time.Hour * time.Duration(interval))
	case "s":
		d1 = dt.Add(time.Second * time.Duration(interval))

	case "m":
		d1 = dt.Add(time.Minute * time.Duration(interval))
	}
	return d1
}

var timeFormat string = "03:04:05 PM"

func main() {

	dt := time.Now()
	fmt.Println("current time", dt.Format(timeFormat))

	d1 := dt.Add(time.Second * 5)

	if len(os.Args) > 2 {
		period := os.Args[1]
		interval, err := strconv.Atoi(os.Args[2])

		if err != nil {
			log.Fatalf(`Please pass interval as args :
				main.go m 10
				main.go s 10
				main.go h 1
			`)
		}
		title := ""
		if len(os.Args) == 4 {
			title = os.Args[3]
		}

		startInterval(d1, period, dt, interval, title)
	}

}

func startInterval(d1 time.Time, period string, dt time.Time, interval int, title string) {
	d1 = addInterval(period, d1, dt, int64(interval))
	fmt.Println("Next Chime will happen on :", d1.Format(timeFormat))
	for {
		current := time.Now()
		if d1.Before(current) {

			// fmt.Println("printing in every 5 secs....", current)
			chime(title)

			d1 = addInterval(period, d1, d1, int64(interval))

			fmt.Println("Next Chime will happen on :", d1.Format(timeFormat))

		}
		time.Sleep(time.Second * 1)
	}
}
