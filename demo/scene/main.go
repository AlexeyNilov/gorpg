package main

import (
	"fmt"

	"github.com/AlexeyNilov/gorpg/system/resource"
)

func main() {
	scene := CreateScene("Light")

	light := resource.Resource{
		Value:    100,
		MaxValue: 100,
	}

	eater := NewLightEater()

	air := resource.Resource{
		Value:    0,
		MaxValue: 100,
	}

	for range 4 {
		fmt.Println(scene)
		fmt.Println(light, eater, air)

		eater.Execute(scene.Description, &light, &air)
		scene.Update()
	}
}
