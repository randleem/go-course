// package main

// import "fmt"

// type Animals interface{
// 	NumberOfLegs() int
// 	Lives() string
// }

// type Dog struct{
// 	name string
// 	breed string
// }

// type Gorilla struct{
// 	name string
// 	numberOfTeeth int
// 	colour string
// }

// func PrintInfo(a Animals) {
// 	fmt.Println("This animal lives in",a.Lives() , "and has", a.NumberOfLegs(), "legs")
// }

// func(d *Dog) Lives() string{
// 	return "with humans"
// }

// func(d *Dog) NumberOfLegs() int{
// 	return 4
// }

// func(g *Gorilla) Lives() string{
// 	return "jungle"
// }

// func(g *Gorilla) NumberOfLegs() int{
// 	return 2
// }

// func main(){

// 	dog := Dog{
// 		name:"Sampson", 
// 		breed:"German Shepherd",
// 	}

// 	gorilla := Gorilla{
// 		name: "Bertie", 
// 		numberOfTeeth: 38,
// 		colour: "grey",
// 	}

// 	PrintInfo(&dog)
// 	PrintInfo(&gorilla)
// }