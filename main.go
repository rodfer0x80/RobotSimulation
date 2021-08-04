package main

import (
	"fmt"

	"github.com/trevalkov/motorRobot/engine"
)

// Board
// If side is a, board is a 2D a*a square
const worldSide = 25
const worldSize = worldSide * worldSide

// Graphics
const robotCell = "*"
const worldCell = "."
const pointCell = "x"
const pointRobotOverlapCell = "o"

// Trail Graphics (dev)
//const trailCell = // "x"
//const trailRobotOverlapCell = //"*"

// Robot
type Robot struct {
	// Robot looks like this
	//    xxx
	//   xxoxx
	//	  xxx

	name string

	totalCells int // 11
	mainPos    int // representes the central cell of the robot as the hitbox "o"

	moveSpeed    int // decimal 2f number x,ab so speed 100 = 1,00
	maxMoveSpeed int
	moveAcc      int  // decimal 2f number x,ab acceleration 100 = 1,00
	moving       bool // moving or stopping
	turnLeft     bool
	turnRight    bool

	//func init()

}

func (robot Robot) move(mainPos int, moveSpeed int, moveAcc int, moving bool, turnLeft bool, turnRight bool) int {
	if turnLeft {
		// nothing
	}

	if turnRight {
		// nothing
	}

	if robot.moving {
		//robot.mainPos = oneDimensionalMove(robot.mainPos, robot.moveSpeed)

		if robot.moveSpeed < robot.maxMoveSpeed {
			robot.moveSpeed += robot.moveAcc
		}

	} else {
		//robot.mainPos = oneDimensionalMove(robot.mainPos, robot.moveSpeed)

		if robot.moveSpeed > 0 {
			robot.moveSpeed -= robot.moveAcc
		}
	}

	return robot.mainPos
}

func initSapien() Robot {
	var sapien = Robot{}

	sapien.name = "Sapien"

	sapien.totalCells = 11
	sapien.mainPos = 30

	sapien.moveSpeed = 0
	sapien.maxMoveSpeed = 5
	sapien.moveAcc = 1
	sapien.moving = false
	sapien.turnLeft = false
	sapien.turnRight = false

	return sapien
}

func (robot Robot) connect() string {
	var connected string = "[*] Robot " + robot.name + " connected"
	return connected
}

// InitWorld starts board state
func initWorld(world *[worldSize]string) *[worldSize]string {

	for i := 0; i < worldSize; i++ {
		world[i] = worldCell
	}
	return world
}

//func tick(board [boardSize] string) [boardSize] string

func outOfWorldRight(column int, rowCells int) bool {
	var flag bool = false

	if column == rowCells-1 {
		flag = true
	}

	return flag
}

func drawWorld(world *[worldSize]string) string {
	var worldString string = ""

	// clearScreen()
	for i := 0; i < worldSize; i += worldSide {
		worldString += "\n"

		for ii := 0; ii < worldSide; ii++ {
			worldString += world[i+ii] + " "
		}
	}

	return worldString
}

func main() {
	world := new([worldSize]string)
	world = initWorld(world)
	var connectEngine string = engine.Connect()
	fmt.Println(connectEngine)

	var sapien = initSapien()
	var connectSapien string = sapien.connect()
	fmt.Println(connectSapien)

	fmt.Println(drawWorld(world))
}
