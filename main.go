package main

import (
	"MyGoPractice/assembler_factory"
	"MyGoPractice/brain_module"
	"MyGoPractice/head_module"
	"MyGoPractice/left_arm_module"
	"MyGoPractice/left_leg_module"
	"MyGoPractice/right_arm_module"
	"MyGoPractice/right_leg_module"
	"MyGoPractice/torso_module"
)

func setupTorso() torso_module.Torso {
	return torso_module.Torso{Model: "T1-Gen1"}
}

func setupLeftLeg() left_leg_module.LeftLeg {
	return left_leg_module.LeftLeg{Model: "L1-L"}
}

func setupRightLeg() right_leg_module.RightLeg {
	return right_leg_module.RightLeg{Model: "L1-R"}
}

func setupLeftArm() left_arm_module.LeftArm {
	return left_arm_module.LeftArm{Model: "A1-L"}
}

func setupRightArm() right_arm_module.RightArm {
	return right_arm_module.RightArm{Model: "A1-R"}
}

func setupHead() head_module.Head {
	return head_module.Head{}
}

func setupBrainMemoryModule() brain_module.BrainMemoryModule {
	return brain_module.BrainMemoryModule {
		Model: "AI-HiveMind v1.0",
		Size: 1000,
		Type: "Random Access Memory 128-GB",
	}
}

func setupBrainPersonalityModule() brain_module.BrianPersonalityModule {
	return brain_module.BrianPersonalityModule {
		Type: "Neutral Negotiator",
		Aggression: 2,
		LogicLevel: 8,
		DiplomacyLevel: 10,
	}
}

func setupBrain() brain_module.Brain {
	return brain_module.Brain{
		Memory: setupBrainMemoryModule(),
		Personality: setupBrainPersonalityModule(),
		Name: "Terminator",
		Version: "v1-26-03-2022",
		Description: "Land Robot",
	}
}

func connectBotModules(botModules assembler_factory.AssembledRobot) assembler_factory.AssembledRobot {
	return assembler_factory.AssembleRobot(botModules)
}

func main() {
	botModules := assembler_factory.AssembledRobot{
		Torso:    setupTorso(),
		LeftArm:  setupLeftArm(),
		RightArm: setupRightArm(),
		LeftLeg:  setupLeftLeg(),
		RightLeg: setupRightLeg(),
		Head: setupHead(),
		Brain:    setupBrain(),
	}
	robot := connectBotModules(botModules)
	robot.Brain.Name = "Terminator v1.2"
	robot.Initialize()
	robot.Identify()
	robot.SelfCheck()
	robot.Calculate(func() int {
		equationSolution := 1.6/1
		return int(equationSolution)
	})
}
