package space

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPoint(t *testing.T) {
	point := NewPoint()
	assert.IsType(t, &Point{}, point)
}

func TestLinkPoints(t *testing.T) {
	point1 := NewPoint()
	point2 := NewPoint()
	LinkPoints(point1, point2)
	assert.Equal(t, point2, point1.Next)
}

func TestAddPoint(t *testing.T) {
	point := NewPoint()
	space := Space{}
	space.AddPoint(point)
	assert.Equal(t, point, &space[0])
	assert.Equal(t, 1, len(space))
}

func TestCreateSpace(t *testing.T) {
	spaceSize := 10
	space := CreateSpace(spaceSize)
	assert.Equal(t, spaceSize, len((*space)))
	for i := range *space {
		assert.NotNil(t, (*space)[i].Next)
	}

	// test if the first and last points are linked
	assert.Equal(t, &(*space)[0], (*space)[len(*space)-1].Next)

	// test ID
	assert.Equal(t, 0, (*space)[0].ID)
}

func BenchmarkCreateSpace(b *testing.B) {
	CreateSpace(10)
}