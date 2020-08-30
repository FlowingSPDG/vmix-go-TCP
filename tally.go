package vmixtcp

// TallyStatus alias to uint
//go:generate stringer -type=TallyStatus
type TallyStatus uint

// TallyResponse TALLY Event response
type TallyResponse struct {
	Status string
	Tally  []TallyStatus
}

const (
	Off TallyStatus = iota
	Program
	Preview
)

// RegisterTallyCallback tally callback for TALLY event.
func (v *Vmix) RegisterTallyCallback(cb func(*TallyResponse)) {
	v.tallyHandler = append(v.tallyHandler, cb)
}

// ClearTallyCallback clear callback for TALLY event.
func (v *Vmix) ClearTallyCallback() {
	v.tallyHandler = make([]func(*TallyResponse), 0, 0)
}
