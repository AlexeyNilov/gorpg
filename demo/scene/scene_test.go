package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateScene(t *testing.T) {
	scene := CreateScene("Light")
	assert.Equal(t, "Light", scene.Description)
}

func TestUpdateScene(t *testing.T) {
	scene := CreateScene("Light")
	scene.Update()
	assert.Equal(t, "Dark", scene.Description)
}
