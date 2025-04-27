package main

import (
	"fmt"

	"github.com/AlexeyNilov/gorpg/ooda"
)

type demoObserver struct{
	State string
}

func (d *demoObserver) Observe() string {
	return d.State
}

type demoOrienter struct{}

func (d *demoOrienter) Orient(observation string) string {
	if observation == "1" {
		return "Good"
	} else {
		return "Bad"
	}
}

type demoDecider struct{}

func (d *demoDecider) Decide(orientation string) string {
	if orientation == "Good" {
		return "Go"
	} else {
		return "Wait"
	}
}

type demoActuator struct{}

func (d *demoActuator) Act(decision string) string {
	if decision == "Go" {
		return "Step"
	} else {
		return ""
	}
}

func main() {
	fmt.Print("OODA loop demo\n")

	process := ooda.OODAProcess{}
	input := demoObserver{"1"}
	process.Run(&input, &demoOrienter{}, &demoDecider{}, &demoActuator{})
	fmt.Print(process, "\n")
	input = demoObserver{"0"}
	process.Run(&input, &demoOrienter{}, &demoDecider{}, &demoActuator{})
	fmt.Print(process, "\n")
}
