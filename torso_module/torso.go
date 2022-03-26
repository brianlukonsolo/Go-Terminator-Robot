package torso_module

import (
	"MyGoPractice/common/strings"
	"MyGoPractice/logging"
)

const ModuleName = "Torso Module"

type Torso struct {
	Model string
	IsActive bool
}

func (t *Torso) SetActive (){
	t.IsActive = true
}

func (t *Torso) SetInactive (){
	t.IsActive = false
}

func (t *Torso) SelfCheck () {
	if t.IsActive == true {
		logging.LOG(ModuleName,strings.Operational)
	}
	if t.IsActive == false {
		logging.LOG(ModuleName,strings.NotOperational)
	}
}