package ooda

type OODAProcess struct {
	Observation string // Data or input observed
	Orientation string // Context or orientation based on observations
	Decision    string // Decision derived from the context
	Action      string // Action planned based on the decision
}

type Observer interface {
	Observe() string
}

type Orienter interface {
	Orient(observation string) string
}

type Decider interface {
	Decide(context string) string
}

// Reset clears all data in the OODA process to restart the loop.
func (o *OODAProcess) Reset() {
	o.Observation = ""
	o.Orientation = ""
	o.Decision = ""
	o.Action = ""
}

func (o *OODAProcess) Observe(observer Observer) {
	o.Observation = observer.Observe()
}

func (o *OODAProcess) Orient(orienter Orienter) {
	// Empty observation leads to empty orientation
	if o.Observation == "" {
		o.Orientation = ""
		return
	}
	o.Orientation = orienter.Orient(o.Observation)
}

func (o *OODAProcess) Decide(decider Decider) {
	// Empty orientation leads to empty decision
	if o.Orientation == "" {
		o.Decision = ""
		return
	}
	o.Decision = decider.Decide(o.Orientation)
}
