package main

import (
	"github.com/AlexeyNilov/gorpg/system/resource"
	"github.com/AlexeyNilov/gorpg/system/transformer"
)

type LightEater struct {
	transformer.ResourceTransformer
}

func CreateLightEater() LightEater {
	le := LightEater{}
	le.MaxValue = 10
	return le
}

func (le *LightEater) Execute(scene Scene, in resource.Getter, out resource.Putter) {
	switch scene.Description {
	case "Light":
		le.Consume(in, 2)
		le.Produce(out, 1)
	}
}
