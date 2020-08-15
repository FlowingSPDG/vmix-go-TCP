package vmixtcp

import (
	"fmt"
)

// Register goroutine callback event
func (v *Vmix) Register(command string, cb func(Response)) error {
	if _, exist := v.subhandler[command]; exist {
		return fmt.Errorf("Handler exist")
	}
	v.subhandler[command] = cb
	return nil
}

// Unregister goroutine callback event
func (v *Vmix) Unregister(command string) error {
	if _, exist := v.subhandler[command]; !exist {
		return fmt.Errorf("Handler not exist")
	}
	delete(v.subhandler, command)
	return nil
}
