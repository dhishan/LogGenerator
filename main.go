package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	counter := 0
	hostname, _ := os.Hostname()
	minute := time.Now().Minute()
	SLEEP_TIME_S := os.Getenv("SLEEP_TIME")
	if SLEEP_TIME_S == "" {
		SLEEP_TIME_S = "5"
	}
	SLEEP_TIME, err := strconv.Atoi(SLEEP_TIME_S)
	if err != nil {
		log.Panic("unable to convert string to int, Error: ", err.Error())
	}
	log.Printf("Sleep Time passed is: %v\n", SLEEP_TIME)
	for {
		log.Printf("Log from %s counter:%v", hostname, counter)
		if time.Now().Minute() != minute {
			counter = 0
			minute = time.Now().Minute()
			log.Printf("changing counter to 0")
		}
		counter++
		time.Sleep(time.Duration(SLEEP_TIME) * time.Second)
	}
}
