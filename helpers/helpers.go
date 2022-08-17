package helpers

import (
	"math/rand"
	"time"
)

type SomeStruct struct {
	TypeName string
	TypeNumber int
}

func RandomNumber(n int) int{
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}