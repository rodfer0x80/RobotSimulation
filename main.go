package main

import (
	"fmt"

	"github.com/trevalkov/motorRobot/engine"
)

func main() {
	var eng string

	fmt.Println("robot")

	eng = engine.Engine("robot")
	fmt.Println(eng)
}
