package main

type Scene struct {
	Description string
}

func CreateScene(description string) Scene {
	return Scene{description}
}

func (s *Scene) Update() {
	switch s.Description {
	case "Light":
		s.Description = "Dark"
	case "Dark":
		s.Description = "Light"
	}
}
