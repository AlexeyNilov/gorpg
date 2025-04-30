package main

import (
	"fmt"

	"github.com/AlexeyNilov/gorpg/system/space"
)

func main() {
	space := space.CreateSpace(10)
	for _, point := range *space {
		fmt.Print(point.ID, "-")
	}

	fmt.Println()
	for _, point := range *space {
		fmt.Print(point.Next.Next.ID, "-")
	}

	fmt.Println()
	point := space.GetPoint(5)
	fmt.Println(point.ID)
}