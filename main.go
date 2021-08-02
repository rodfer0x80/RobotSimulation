package main

import (
	"fmt"

	"github.com/trevalkov/motorRobot/engine"
)

func main() {
	var conEng string
	conEng = engine.Connect()
	fmt.Println(conEng)

	fmt.Println("robot")
}
