package vmixtcp

// TallyStatus alias to int
type TallyStatus uint

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
