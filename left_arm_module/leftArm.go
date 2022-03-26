package left_arm_module

import (
	"MyGoPractice/common/strings"
	"MyGoPractice/logging"
)

const ModuleName = "Left Arm Module"

type LeftArm struct {
	Model string
	IsActive bool
}

func (la *LeftArm) SetActive (){
	la.IsActive = true
}

func (la *LeftArm) SetInactive (){
	la.IsActive = false
}

func (la *LeftArm) SelfCheck () {
	if la.IsActive == true {
		logging.LOG(ModuleName,strings.Operational)
	}
	if la.IsActive == false {
		logging.LOG(ModuleName,strings.NotOperational)
	}
}