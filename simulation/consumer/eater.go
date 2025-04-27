package main

import (
	"fmt"

	"github.com/AlexeyNilov/gorpg/ooda"
	"github.com/AlexeyNilov/gorpg/system/transformer"
)

type Eater struct {
	transformer.Transformer
	PointOfView string
}

func (e *Eater) Observe() string {
	return e.PointOfView
}

func (e *Eater) Orient(observation string) string {
	if observation == "Day" {
		return "Time to eat"
	}
	return "Time to sleep"
}

func (e *Eater) Decide(orientation string) string {
	if orientation == "Time to eat" {
		return "Eat"
	}
	return "Sleep"
}

func (e *Eater) Act(decision string) string {
	if decision == "Eat" {
		return "Consume"
	}
	return ""
}

func (e *Eater) Execute(process ooda.OODAProcess) {
	process.Run(e, e, e, e)
	if process.Action == "Consume" {
		e.Transformer.Consume()
	}
}

func main() {
	process := ooda.OODAProcess{}
	eater := Eater{Transformer: transformer.Transformer{}}
	eater.PointOfView = "Day"
	eater.Execute(process)
	eater.Execute(process)
	eater.PointOfView = "Night"
	eater.Execute(process)
	fmt.Print(eater, "\n")
}
