package main

type Robot struct {
	name string

	mainPos int
	moving  bool

	moveSpeed     int
	maxMoveSpeed  int
	moveAcc       int
	turningWheels bool
}

func (robot Robot) connect() string {
	var connected string = "[*] Robot " + robot.name + " connected"
	return connected
}

func (robot Robot) haltingTime(n int) int {
	var totalTime int

	for i := 1; i <= n; i++ {
		totalTime += i
	}

	return totalTime
}

func (robot Robot) resetTurningWheels() bool {
	return true
}

func (robot Robot) move(mark int) (int, int, bool, bool) {
	var horizontalMove int
	var verticalMove int

	horizontalMove, verticalMove, robot.moveSpeed, robot.turningWheels = robot.controlWheels(mark)

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
	}

	return robot.mainPos, robot.moveSpeed, robot.moving, robot.turningWheels
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
		if robot.turningWheels {
			robot.turningWheels = false
			robot.moveSpeed -= robot.moveAcc
		}

		if abs(verticalMove) >= robot.haltingTime(robot.moveSpeed+robot.moveAcc) {
			robot.moveSpeed += robot.moveAcc
		} else if abs(verticalMove) >= robot.haltingTime(robot.moveSpeed) {
		} else {
			robot.moveSpeed -= robot.moveAcc
		}
	}

	return horizontalMove, verticalMove, robot.moveSpeed, robot.turningWheels
}
