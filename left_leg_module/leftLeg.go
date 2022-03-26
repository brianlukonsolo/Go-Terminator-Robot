package left_leg_module

import (
	"MyGoPractice/common/strings"
	"MyGoPractice/logging"
)

const ModuleName = "Left Leg Module"

type LeftLeg struct {
	Model string
	IsActive bool
}

func (ll *LeftLeg) SetActive (){
	ll.IsActive = true
}

func (ll *LeftLeg) SetInactive (){
	ll.IsActive = false
}

func (ll *LeftLeg) SelfCheck () {
	if ll.IsActive == true {
		logging.LOG(ModuleName,strings.Operational)
	}
	if ll.IsActive == false {
		logging.LOG(ModuleName,strings.NotOperational)
	}
}