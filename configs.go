package main

// World
const worldLateral = 30

// Graphics
const worldCell = "."
const robotCell = "*"
const markCell = "x"
const robotMarkOverlapCell = "o"

// Robot
func initRobot() Robot {
	var robot = Robot{}

	robot.name = "Sapien"

	robot.mainPos = worldSize/2 + worldLateral/2
	robot.moving = false

	robot.moveSpeed = 0
	robot.maxMoveSpeed = 5
	robot.moveAcc = 1
	robot.turningWheels = true

	return robot
}
