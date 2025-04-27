package main

import (
	"fmt"

	"github.com/AlexeyNilov/gorpg/ooda"
)

type demoObject struct{
	State string
}

func (d *demoObject) Observe() string {
	return d.State
}

func (d *demoObject) Orient(observation string) string {
	if observation == "1" {
		return "Good"
	} else {
		return "Bad"
	}
}

func (d *demoObject) Decide(orientation string) string {
	if orientation == "Good" {
		return "Go"
	} else {
		return "Wait"
	}
}

func (d *demoObject) Act(decision string) string {
	if decision == "Go" {
		return "Step"
	} else {
		return ""
	}
}

func (d *demoObject) Execute(process ooda.OODAProcess) {
	process.Run(d, d, d, d)
	fmt.Print(process, "\n")
	if process.Action != "" {
		fmt.Print("Do ", process.Action, "\n")
	}
}

func main() {
	fmt.Print("OODA loop demo\n")

	process := ooda.OODAProcess{}
	obj := demoObject{State: "1"}
	obj.Execute(process)

	obj.State = "0"
	obj.Execute(process)
}
