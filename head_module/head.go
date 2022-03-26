package head_module

import (
	"MyGoPractice/common/strings"
	"MyGoPractice/logging"
)

const ModuleName = "Head Module"

type Head struct {
	Model string
	IsActive bool
}

func (h *Head) SetActive (){
	h.IsActive = true
}

func (h *Head) SetInactive (){
	h.IsActive = false
}

func (h *Head) SelfCheck () {
	if h.IsActive == true {
		logging.LOG(ModuleName,strings.Operational)
	}
	if h.IsActive == false {
		logging.LOG(ModuleName,strings.NotOperational)
	}
}