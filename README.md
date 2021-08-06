# motorRobot-engine
2D engine and motor robot simulation in Golang

### Design
##### Engine
2D matrix in a 1D array with border and position checking and two points A and B where the robot should start and finish respectively asserted randomly in each hemisphere 
##### Robot
Robot simulator with wheels that moves [up, down], slowly turning [right, left] (possible while moving), with direction switching with [start/stop] on [spacebar] with negative acceleration to stop, acceleration, speed [0-x] and distance sensoring for acceleration management
