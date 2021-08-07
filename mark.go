package main

import (
	"math/rand"
	"time"
)

func placeMark() int {
	rand.Seed(time.Now().UnixNano())
	var randomNumber int = rand.Intn(worldSize)

	return randomNumber
}
