package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/toast.v1"
)

func chime() {

	notification := toast.Notification{
		AppID:   "Chime",
		Title:   "Take deep breath and move",
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

func main() {

	dt := time.Now()
	fmt.Println("current time", dt.String())

}
