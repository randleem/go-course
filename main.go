package main

import "fmt"

var nyName string

func main(){
	fmt.Println("Hello Emma")

	var whatToSay string
	var i int

	whatToSay = "Goodbye cruel world"
	fmt.Println(whatToSay)

	i = 7
	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThing := saySomething()
	fmt.Println("The function returned", whatWasSaid, theOtherThing)
}

func saySomething() (string, string){
	return "something", "else"
}