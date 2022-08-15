package main

import "log"

func main(){
	 var myString string 
	 myString = "Green"

	 log.Println("myString is set to", myString)
	 changeUsingPointer(&myString)
	//  reference to a variable add &
	 log.Println("after func call myString is set to", myString)
}

func changeUsingPointer(s *string){
	// refer to a pointer add * - pointer allows you to change a variable in a different scope by getting where its stored on the computers memory
	newValue := "Red"
	*s = newValue
}