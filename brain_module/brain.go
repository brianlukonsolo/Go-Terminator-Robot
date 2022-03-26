package brain_module

import (
	"MyGoPractice/common/strings"
	"MyGoPractice/logging"
)

const ModuleName = "Brain Module"


type BrainMemoryModule struct {
	Model string
	IsActive bool
	Size int
	Type string
}

type BrianPersonalityModule struct {
	Type string
	Aggression int
	LogicLevel int
	DiplomacyLevel int
}

type Brain struct {
	Memory BrainMemoryModule
	Personality BrianPersonalityModule
	Name string
	Version string
	Description string
	IsActive bool
}

func (b *Brain) SetActive (){
	b.IsActive = true
}

func (b *Brain) SetInactive (){
	b.IsActive = false
}

func (b *Brain) SelfCheck () {
	if b.IsActive == true {
		logging.LOG(ModuleName, strings.Operational)
	}
	if b.IsActive == false {
		logging.LOG(ModuleName, strings.NotOperational)
	}
}
