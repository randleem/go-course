package main

import "log"

type myStruct struct{
	FirstName string
}

func (m *myStruct) printFirstName() string{
	return m.FirstName
}

func main(){
	var myVar myStruct
	myVar.FirstName = "Emma"
	myVar2:= myStruct{
		FirstName: "Bob",
	}
	log.Println("my var is set to", myVar.printFirstName())
	log.Println("my var is set to", myVar2.printFirstName())
}