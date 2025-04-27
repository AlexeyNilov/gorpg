package ooda

type OODAProcess struct {
	Observation string // Data or input observed
	Orientation string // Orientation(context) based on observations
	Decision    string // Decision(plan, strategy) derived from the context
	Action      string // Action based on the decision
}

type Observer interface {
	Observe() string
}

type Orienter interface {
	Orient(observation string) string
}

type Decider interface {
	Decide(orientation string) string
}

type Actuator interface {
	Act(decision string) string
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

func (o *OODAProcess) Act(actuator Actuator) {
	// Empty decision leads to empty action
	if o.Decision == "" {
		o.Action = ""
		return
	}
	o.Action = actuator.Act(o.Decision)
}

func (o *OODAProcess) Run(observer Observer, orienter Orienter, decider Decider, actuator Actuator) {
	o.Reset()
	o.Observe(observer)
	o.Orient(orienter)
	o.Decide(decider)
	o.Act(actuator)
}
