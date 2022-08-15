package main

import (
	"log"
	"time"
)

type User struct{
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	Birthdate time.Time
	}

func main(){
	
	user := User{
		FirstName: "Trevor", LastName: "Sawler", PhoneNumber:"01789721479", Age:34,
	}
	log.Println(user.PhoneNumber)
	
}
