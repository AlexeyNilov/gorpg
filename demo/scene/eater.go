package main

import (
	"github.com/AlexeyNilov/gorpg/system/resource"
	"github.com/AlexeyNilov/gorpg/system/transformer"
)

type LightEater struct {
	transformer.ResourceTransformer
	ConsumeRate int
	ProduceRate int
}

func NewLightEater() LightEater {
	le := LightEater{}
	le.ConsumeRate = 2
	le.ProduceRate = 1
	le.MaxValue = 10
	return le
}

func (le *LightEater) Execute(sceneDescription string, in resource.Getter, out resource.Putter) {
	switch sceneDescription {
	case "Light":
		le.Consume(in, le.ConsumeRate)
		le.Produce(out, le.ProduceRate)
	}
}
