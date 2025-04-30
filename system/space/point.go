package space

type Point struct {
	Next *Point
	ID int
}

type Space []Point

func NewPoint() *Point {
	return &Point{}
}

func LinkPoints(p1, p2 *Point) {
	p1.Next = p2
}

func (s *Space) AddPoint(p *Point) {
	*s = append(*s, *p)
}

func CreateSpace(n int) *Space {
	space := make(Space, n) // Pre-allocate space for n points
	for i := range n {
		space[i] = *NewPoint() // Initialize each point in the space
		space[i].ID = i
		LinkPoints(&space[i], &space[(i+1)%n])
	}

	return &space
}

func (s *Space) GetPoint(id int) *Point {
	for i := range *s {
		if (*s)[i].ID == id {
			return &(*s)[i]
		}
	}
	return nil
}


