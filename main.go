package main

import (
	"log"
	"time"
)

func main() {

	// The date string is on time.RubyDate format
	_, err := time.Parse(time.UnixDate, "Fri Jan 01 03:01:00 +0000 2016")
	if err == nil {
		log.Fatal("no error")
	}

}
