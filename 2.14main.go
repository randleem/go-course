package main

import "fmt"

func main(){
	// for i :=0; i<=10;i++{
	// 	fmt.Println(i)
	// } 

	// animals := []string{"dog", "fish", "horse", "cow", "cat"}
	// for _, animal := range animals{
	// 	fmt.Println(animal)
	// }

	// animals := make(map[string]string)
	// animals["dog"] = "Fido"
	// animals["cat"] = "Fluffy"
	// for animalType, animal := range animals{
	// 		fmt.Println(animalType, animal)
	// 	}

	// firstLine := "Once upon a time"
	// for i, l := range firstLine{
	// 	fmt.Println(i, l)
	// }

	type User struct{
		firstName string
		lastName string
		email string
		age int
	}

	var users []User
	users = append(users, User{"Emma", "Randle", "fhgheih", 34})
	users = append(users, User{"Dan", "Hayward", "fhgheih", 37})
	users = append(users, User{"Reggie", "Randle", "fhgheih", 0})

	for _, user := range users{
		fmt.Println(user.age)
	}
}