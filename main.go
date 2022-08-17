package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the homepage handler
func Home(w http.ResponseWriter, r *http.Request){
	_, _ = fmt.Fprintf(w, "This is the homepage")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request){
	sum := addValues(4, 5)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("this is the about page, 4+5 is %d", sum))
}

// Divide is the Divide page handler
func Divide(w http.ResponseWriter, r *http.Request){
	f, err := divide(4.0, 5.0)
	if err != nil{
		fmt.Fprintf(w, "cannot divide by zero")
		return
	}
	_, _ = fmt.Fprintf(w, fmt.Sprintf("this is the divide page, %f / %f is %f",4.0, 5.0, f))
}

func divide(x, y float32) (float32, error){
	if y <=0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x/y
	return result, nil
}

//addvalues adds to integers and returns their sum
func addValues(x, y int) (int){
	return x+y
}

func main(){
	
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}