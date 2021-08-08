package main

func getCleanScreen() string {
	var cleanScreen string = "\033[2J"
	return cleanScreen
}

func initWorld(world *[worldSize]string) *[worldSize]string {
	for i := 0; i < worldSize; i++ {
		world[i] = worldCell
	}

	return world
}

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

