package main

import (
	"fmt"

	"github.com/trevalkov/motorRobot/engine"
)

// Board
// If side is a, board is a 2D a*a square
const boardSide = 400
const boardSize = boardSide * boardSide

// Graphics
const robotCell = "*"
const emptyCell = "."
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
	totalCells int // 11
	mainCell   int // representes the central cell of the robot as the hitbox "o"

	speed               int  // decimal 2f number x,ab so speed 100 = 1,00
	acceleration        int  // decimal 2f number x,ab acceleration 100 = 1,00
	accelerationBalance bool // negative or positive acceleration true/false
	status              bool // moving or stopping

}

// InitBoard starts board state
func initBoard() string {
	var board = new([boardSize]string)

	return board[0]
}

//func tick(board [boardSize] string) [boardSize] string

func outOfBoardRight(column int, rowCells int) bool {
	var flag bool = false

	if column == rowCells-1 {
		flag = true
	}

	return flag
}

func main() {
	var conEng string
	conEng = engine.Connect()
	fmt.Println(conEng)

	fmt.Println("robot")
}
