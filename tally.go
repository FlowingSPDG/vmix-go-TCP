package vmixtcp

// TallyStatus alias to uint
type TallyStatus uint

func (t TallyStatus) String() string {
	switch t {
	case TallyOff:
		return "Off"
	case TallyProgram:
		return "Program"
	case TallyPreview:
		return "Preview"
	default:
		return "UNKNOWN"
	}
}

// TallyResponse TALLY Event response
type TallyResponse struct {
	Status string
	Tally  []TallyStatus
}

const (
	// TallyOff = 0
	TallyOff TallyStatus = iota
	// TallyProgram = 1
	TallyProgram
	// TallyPreview = 2
	TallyPreview
)

// RegisterTallyCallback tally callback for TALLY event.
func (v *Vmix) RegisterTallyCallback(cb func(*TallyResponse)) {
	v.tallyHandler = append(v.tallyHandler, cb)
}

// ClearTallyCallback clear callback for TALLY event.
func (v *Vmix) ClearTallyCallback() {
	v.tallyHandler = make([]func(*TallyResponse), 0, 0)
}
