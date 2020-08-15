package vmixtcp

import (
	"fmt"
)

// Register goroutine callback event. Use tally event if it's for TALLY!!
func (v *Vmix) Register(command string, cb func(*Response)) error {
	if _, exist := v.cbhandler[command]; exist {
		return fmt.Errorf("Handler exist")
	}
	v.cbhandler[command] = cb
	return nil
}

// Unregister goroutine callback event
func (v *Vmix) Unregister(command string) error {
	if _, exist := v.cbhandler[command]; !exist {
		return fmt.Errorf("Handler not exist")
	}
	delete(v.cbhandler, command)
	return nil
}
