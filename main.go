package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/trevalkov/motorRobot/engine"
)

// World
// When side equals a, board is a 2D a*a square
const worldLateral = 30
const worldSize = worldLateral * worldLateral

// Graphics
const worldCell = "."
const robotCell = "*"
const markCell = "x"
const robotMarkOverlapCell = "o"

// Robot
type Robot struct {
	name string

	mainPos int // representes the central cell of the robot as the hitbox "o"

	moveSpeed    int // decimal 2f number x,ab so speed 100 = 1,00
	maxMoveSpeed int
	moveAcc      int // decimal 2f number x,ab acceleration 100 = 1,00

	moving  bool // moving or stopping
	turning bool
}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func (robot Robot) move(mark int) (int, int, bool, bool) {
	var horizontalMove int
	var verticalMove int

	horizontalMove, verticalMove, robot.moveSpeed, robot.turning = robot.controlWheels(mark)

	if robot.moveSpeed > 0 {
		robot.moving = true
	} else {
		robot.moving = false
	}

	if robot.moving {
		if horizontalMove > 0 {
			robot.mainPos += robot.moveSpeed
		} else if horizontalMove < 0 {
			robot.mainPos -= robot.moveSpeed
		} else {
			if verticalMove > 0 {
				robot.mainPos += robot.moveSpeed * worldLateral
			} else if verticalMove < 0 {
				robot.mainPos -= robot.moveSpeed * worldLateral
			} else {
				robot.moving = false
			}
		}
	} else {
		// robot stopped
	}

	return robot.mainPos, robot.moveSpeed, robot.moving, robot.turning
}

func (robot Robot) resetTurningWheels() bool {
	return true
}
func initSapien() Robot {
	var sapien = Robot{}

	sapien.name = "Sapien"

	sapien.mainPos = worldSize/2 + worldLateral/2

	sapien.moveSpeed = 0
	sapien.maxMoveSpeed = 5
	sapien.moveAcc = 1

	sapien.moving = false
	sapien.turning = true

	return sapien
}

func (robot Robot) haltingTime(n int) int {
	var totalTime int

	for i := 1; i <= n; i++ {
		totalTime += i
	}

	return totalTime
}

func (robot Robot) connect() string {
	var connected string = "[*] Robot " + robot.name + " connected"
	return connected
}

func (robot Robot) controlWheels(mark int) (int, int, int, bool) {
	var markRow int = mark / worldLateral
	var robotRow int = robot.mainPos / worldLateral

	var markColumn int = mark - (worldLateral * markRow)
	var robotColumn int = robot.mainPos - (worldLateral * robotRow)

	var verticalMove int = markRow - robotRow
	var horizontalMove int = markColumn - robotColumn

	if abs(horizontalMove) > 0 {
		if abs(horizontalMove) >= robot.haltingTime(robot.moveSpeed+robot.moveAcc) {
			robot.moveSpeed += robot.moveAcc
		} else if abs(horizontalMove) >= robot.haltingTime(robot.moveSpeed) {
		} else {
			robot.moveSpeed -= robot.moveAcc
		}
	} else {
		if abs(verticalMove) >= robot.haltingTime(robot.moveSpeed+robot.moveAcc) {
			if !(robot.turning) {
				robot.moveSpeed += robot.moveAcc
			} else {
				robot.turning = false
			}
		} else if abs(verticalMove) >= robot.haltingTime(robot.moveSpeed) {
		} else {
			robot.moveSpeed -= robot.moveAcc
		}
	}

	return horizontalMove, verticalMove, robot.moveSpeed, robot.turning
}

func placeMark() int {
	rand.Seed(time.Now().UnixNano())
	var randomNumber int = rand.Intn(worldSize)

	return randomNumber
}

func getCleanScreen() string {
	var cleanScreen string = "\033[2J"
	return cleanScreen
}

// InitWorld starts board state
func initWorld(world *[worldSize]string) *[worldSize]string {
	for i := 0; i < worldSize; i++ {
		world[i] = worldCell
	}

	return world
}

// InitWorld starts board state
func drawWorld(world *[worldSize]string, mark int, robot Robot) *[worldSize]string {
	for i := 0; i < worldSize; i++ {
		if i == mark && i == robot.mainPos {
			world[i] = robotMarkOverlapCell
		} else if i == robot.mainPos {
			world[i] = robotCell
		} else if i == mark {
			world[i] = markCell
		} else {
			world[i] = worldCell
		}
	}
	return world
}

func getWorld(world *[worldSize]string) string {
	var worldString string = ""

	for i := 0; i < worldSize; i += worldLateral {
		worldString += "\n"

		for ii := 0; ii < worldLateral; ii++ {
			worldString += world[i+ii] + " "
		}
	}

	return worldString
}

func tick(robot Robot, world *[worldSize]string) {
	var mark int = placeMark()
	robot.turning = robot.resetTurningWheels()

	for 1 > 0 {
		time.Sleep(1 * time.Second)
		fmt.Println(getCleanScreen())

		robot.mainPos, robot.moveSpeed, robot.moving, robot.turning = robot.move(mark)

		world = drawWorld(world, mark, robot)
		fmt.Println(getWorld(world))

		if robot.mainPos == mark {
			fmt.Println("[*] Resetting...")
			time.Sleep(3 * time.Second)

			tick(robot, world)
		}
	}
}

func main() {
	world := new([worldSize]string)
	world = initWorld(world)

	var connectEngine string = engine.Connect()
	fmt.Println(connectEngine)

	var sapien = initSapien()
	var connectSapien string = sapien.connect()
	fmt.Println(connectSapien)

	tick(sapien, world)
}
