package main

import (
	"fmt"

	"github.com/AlexeyNilov/gorpg/system/resource"
	"github.com/AlexeyNilov/gorpg/system/transformer"
)

func main() {

	Light := resource.Resource{
		Value:    100,
		MaxValue: 100,
	}
	Tree := transformer.ResourceTransformer{
		Resource: resource.Resource{
			Value:    0,
			MaxValue: 10,
		},
	}
	Oxygen := resource.Resource{
		Value:    0,
		MaxValue: 100,
	}

	for range 10 {
		Tree.Consume(&Light, 1)
		Tree.Produce(&Oxygen, 1)
	}
	
	fmt.Println(Light, Tree, Oxygen)
}
