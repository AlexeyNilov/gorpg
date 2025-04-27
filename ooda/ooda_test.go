package ooda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TestObservation = "Test Observation"
	TestOrientation = "Test Orientation"
	TestDecision    = "Test Decision"
	TestAction      = "Test Action"
)

// newTestOODAProcess initializes an OODAProcess with test data.
func newTestOODAProcess() OODAProcess {
	return OODAProcess{
		Observation: TestObservation,
		Orientation: TestOrientation,
		Decision:    TestDecision,
		Action:      TestAction,
	}
}

func TestReset(t *testing.T) {
	process := newTestOODAProcess()
	process.Reset()

	expected := OODAProcess{
		Observation: "",
		Orientation: "",
		Decision:    "",
		Action:      "",
	}
	assert.Equal(t, expected, process)
}

type MockObserver struct {
	Data string
}

func (m MockObserver) Observe() string {
	return m.Data
}

type MockOrienter struct {
	Data string
}

func (m MockOrienter) Orient(observation string) string {
	return m.Data
}

type MockDecider struct {
	Data string
}

func (m MockDecider) Decide(orientation string) string {
	return m.Data
}

type MockActuator struct {
	Data string
}

func (m MockActuator) Act(decision string) string {
	return m.Data
}

func TestObserve(t *testing.T) {
	mockObserver := MockObserver{Data: TestObservation}
	process := newTestOODAProcess()
	process.Observe(mockObserver)

	assert.Equal(t, TestObservation, process.Observation)
}

func TestOrient(t *testing.T) {
	mockOrienter := MockOrienter{Data: TestOrientation}
	process := newTestOODAProcess()
	process.Orient(mockOrienter)

	assert.Equal(t, TestOrientation, process.Orientation)

	process.Observation = "" // Reset Observation for next test
	process.Orient(mockOrienter)
	assert.Equal(t, "", process.Orientation)
}

func TestDecide(t *testing.T) {
	mockDecider := MockDecider{Data: TestDecision}
	process := newTestOODAProcess()
	process.Decide(mockDecider)

	assert.Equal(t, TestDecision, process.Decision)

	process.Orientation = "" // Reset Orientation for next test
	process.Decide(mockDecider)

	assert.Equal(t, "", process.Decision)
}

func TestAct(t *testing.T) {
	mockActuator := MockActuator{Data: TestAction}
	process := newTestOODAProcess()
	process.Act(mockActuator)

	assert.Equal(t, TestAction, process.Action)

	process.Decision = "" // Reset Decision for next test
	process.Act(mockActuator)

	assert.Equal(t, "", process.Action)
}

func TestWholeProcess(t *testing.T) {
	process := newTestOODAProcess()

	mockObserver := MockObserver{Data: TestObservation}
	process.Observe(mockObserver)

	mockOrienter := MockOrienter{Data: process.Orientation}
	process.Orient(mockOrienter)

	mockDecider := MockDecider{Data: process.Decision}
	process.Decide(mockDecider)

	mockActuator := MockActuator{Data: process.Action}
	process.Act(mockActuator)

	assert.Equal(t, TestAction, process.Action)
}

func TestRun(t *testing.T) {
	process := newTestOODAProcess()

	mockObserver := MockObserver{Data: TestObservation}
	mockOrienter := MockOrienter{Data: TestOrientation}
	mockDecider := MockDecider{Data: TestDecision}
	mockActuator := MockActuator{Data: TestAction}

	process.Run(mockObserver, mockOrienter, mockDecider, mockActuator)

	assert.Equal(t, TestAction, process.Action)
}


