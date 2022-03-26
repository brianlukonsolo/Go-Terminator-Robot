package common

type Robot interface {
	Initialize()
	Identify()
	SelfCheck()
}

type RobotPart interface {
	SetActive()
	SetInactive()
	SelfCheck()
}
