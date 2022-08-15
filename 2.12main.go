package main

import (
	"fmt"
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName string
}

func main(){

	myMap := make(map[string] string)

	myMap["dog"] = "Samson"
	myMap["otherDog"] = "Cassie"
	log.Println(myMap["dog"])
	log.Println(myMap["otherDog"])

	myMap["dog"] = "Reggie"
	log.Println(myMap["dog"])

	myOther := make(map[string]int)
	myOther["first"] = 1
	myOther["second"] = 2
	log.Println(myOther["first"], myOther["second"])

	myUser := make(map[string]User)
	me := User{
		FirstName: "Emma",
		LastName: "Randle",
	}

	myUser["me"] = me

	log.Println(myUser)

	var myNewVar float32
	myNewVar = 11.2
	log.Println(myNewVar)

	var mySlice []int
	mySlice = append(mySlice, 4)
	mySlice = append(mySlice,2)
	mySlice = append(mySlice, 3)

	sort.Ints(mySlice)
	
	log.Println(mySlice)

	numbers := []int{1,3,5,7,9,11,13,15,17,19,21}
	log.Println(numbers[4:7])

	names := []string{"one", "frog", "goat"}
	fmt.Println(names)
}