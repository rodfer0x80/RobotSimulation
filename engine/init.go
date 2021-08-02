package engine


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
	totalCells int = 11
	mainCell int // representes the central cell of the robot as the hitbox "o"

	speed int // decimal 2f number x,ab so speed 100 = 1,00
	acceleration int // decimal 2f number x,ab acceleration 100 = 1,00
	accelerationBalance bool // negative or positive acceleration true/false
	status bool // moving or stopping
	
}

// Connect alerts module import
func Connect() string {
	var status string = "[*] Engine connected"
	return status
}

// InitBoard starts board state
func InitBoard() string {
	var board = new([boardSize] string)

}

func Tick(board [boardSize] string) [boardSize] string
func outOfBoardRight(column int, rowCells int) bool {
	var flag bool = false

	if column == rowCells-1 {
		flag = true
	}

	return flag
}

func outOfBoardLeft(column int) bool {
	var flag bool = false

	if column == 0 {
		flag = true
	}

	return flag
}

func outOfBoardUp(row int) bool {
	var flag bool = false

	if row == 0 {
		flag = true
	}

	return flag
}

func outOfBoardDown(row int, rowCells int) bool {
	var flag bool = false

	if row == rowCells-1 {
		flag = true
	}

	return flag
}
