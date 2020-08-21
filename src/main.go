package main

import (
	"log"
	"time"
)

func main() {
	log.Println("Let's GO on an adventure...")
	txt1 := make(chan string)
	txt3 := make(chan string)
	txt10 := make(chan string)
	defer close(txt1)
	defer close(txt3)
	defer close(txt10)

	go remaining1(txt1)
	go remaining3(txt3)
	go remaining10(txt10)

	loop:
		for {
			select {
			case v := <- txt1:
				log.Println("arrived: " + v)
			case v := <- txt3:
				log.Println("arrived: " + v)
			case v := <- txt10:
				log.Println("arrived: " + v)
				break loop
			}
		}
}

func remaining1(txt chan string) {
	for {
		time.Sleep(1 * time.Second)
		txt <- "1 second"
	}
}

func remaining3(txt chan string) {
	for {
		time.Sleep(3 * time.Second)
		txt <- "3 seconds"
	}
}

func remaining10(txt chan string) {
	for {
		time.Sleep(10 * time.Second)
		txt <- "10 seconds"
	}
}