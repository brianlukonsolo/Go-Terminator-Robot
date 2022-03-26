package right_leg_module

import (
	"MyGoPractice/common/strings"
	"MyGoPractice/logging"
)

const ModuleName = "Right Leg Module"

type RightLeg struct {
	Model string
	IsActive bool
}

func (rl *RightLeg) SetActive (){
	rl.IsActive = true
}

func (rl *RightLeg) SetInactive (){
	rl.IsActive = false
}

func (rl *RightLeg) SelfCheck () {
	if rl.IsActive == true {
		logging.LOG(ModuleName,strings.Operational)
	}
	if rl.IsActive == false {
		logging.LOG(ModuleName,strings.NotOperational)
	}
}