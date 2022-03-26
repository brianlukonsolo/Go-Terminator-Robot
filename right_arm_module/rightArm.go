package right_arm_module

import (
	"MyGoPractice/common/strings"
	"MyGoPractice/logging"
)

const ModuleName = "Right Arm Module"

type RightArm struct {
	Model string
	IsActive bool
}

func (ra *RightArm) SetActive (){
	ra.IsActive = true
}

func (ra *RightArm) SetInactive (){
	ra.IsActive = false
}

func (ra *RightArm) SelfCheck () {
	if ra.IsActive == true {
		logging.LOG(ModuleName,strings.Operational)
	}
	if ra.IsActive == false {
		logging.LOG(ModuleName,strings.NotOperational)
	}
}