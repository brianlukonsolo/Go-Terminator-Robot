package assembler_factory

import (
	"MyGoPractice/brain_module"
	"MyGoPractice/head_module"
	"MyGoPractice/left_arm_module"
	"MyGoPractice/left_leg_module"
	"MyGoPractice/logging"
	"MyGoPractice/right_arm_module"
	"MyGoPractice/right_leg_module"
	"MyGoPractice/torso_module"
	"strconv"
)

const ModuleName = "Assembler"

type AssembledRobot struct{
	Brain brain_module.Brain
	Head head_module.Head
	LeftArm left_arm_module.LeftArm
	RightArm right_arm_module.RightArm
	Torso torso_module.Torso
	LeftLeg left_leg_module.LeftLeg
	RightLeg right_leg_module.RightLeg
	Name string
}

func (a AssembledRobot) Initialize() (){
	logging.LOG("Robot", "Initializing ...")
}

func (a AssembledRobot) Identify () {
	logging.LOG(
		"Robot",
		a.Name + " is brain type " + a.Brain.Personality.Type + " version " + a.Brain.Version,
	)
}

func (a *AssembledRobot) SelfCheck () {
	a.LeftLeg.SelfCheck()
	a.RightLeg.SelfCheck()
	a.Torso.SelfCheck()
	a.LeftArm.SelfCheck()
	a.RightArm.SelfCheck()
	a.Head.SelfCheck()
	a.Brain.SelfCheck()
}

// Calculate - accepts any mathematical implementation that returns an int
func (a *AssembledRobot) Calculate (anonymousFunction func() int) int {
	solution := anonymousFunction()
	logging.LOG("Robot", "solution found: " + strconv.Itoa(solution))
	return solution
}

// Private functions
func attachLeftArm (arm left_arm_module.LeftArm, bot AssembledRobot) AssembledRobot {
	bot.LeftArm = arm
	bot.LeftArm.IsActive = true
	logging.LOG(ModuleName, "attached left arm")
	return bot
}

func attachRightArm (arm right_arm_module.RightArm, bot AssembledRobot) AssembledRobot {
	bot.RightArm = arm
	bot.RightArm.IsActive = true
	logging.LOG(ModuleName, "attached right arm")
	return bot
}

func attachRightLeg (leg right_leg_module.RightLeg, bot AssembledRobot) AssembledRobot {
	bot.RightLeg = leg
	bot.RightLeg.IsActive = true
	logging.LOG(ModuleName, "attached right leg")
	return bot
}

func attachLeftLeg (leg left_leg_module.LeftLeg, bot AssembledRobot) AssembledRobot {
	bot.LeftLeg = leg
	bot.LeftLeg.IsActive = true
	logging.LOG(ModuleName, "attached left leg")
	return bot
}

func attachTorso (torso torso_module.Torso, bot AssembledRobot) AssembledRobot {
	bot.Torso = torso
	bot.Torso.IsActive = true
	logging.LOG(ModuleName, "attached torso")
	return bot
}

func attachHead (head head_module.Head, bot AssembledRobot) AssembledRobot {
	bot.Head = head
	bot.Head.SetActive()
	logging.LOG(ModuleName, "attached head")
	return bot
}

func attachBrain (brain brain_module.Brain, bot AssembledRobot) AssembledRobot {
	bot.Brain = brain
	bot.Brain.SetActive()
	logging.LOG(ModuleName, "attached brain")
	return bot
}

func AssembleRobot (modules AssembledRobot) AssembledRobot {
	assembledRobot := AssembledRobot {}
	assembledRobot = attachTorso(modules.Torso, assembledRobot)
	assembledRobot = attachLeftArm(modules.LeftArm, assembledRobot)
	assembledRobot = attachRightArm(modules.RightArm, assembledRobot)
	assembledRobot = attachLeftLeg(modules.LeftLeg, assembledRobot)
	assembledRobot = attachRightLeg(modules.RightLeg, assembledRobot)
	assembledRobot = attachHead(modules.Head, assembledRobot)
	assembledRobot = attachBrain(modules.Brain, assembledRobot)
	return assembledRobot
}