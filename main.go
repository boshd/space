package main

import (
	"fmt"
	"log"
)

func main() {

	// create db connection
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// will migrate User struct to database and will check for
	// changes and will automatically migrate them everytime
	// the service is started.

	fmt.Print("jel")
}
