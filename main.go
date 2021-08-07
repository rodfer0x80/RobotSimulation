package main

import (
	"fmt"
)

const worldSize = worldLateral * worldLateral

func main() {
	world := new([worldSize]string)
	world = initWorld(world)

	var robot = initRobot()
	var connectRobot string = robot.connect()
	fmt.Println(connectRobot)

	tick(robot, world)
}
