package main

import (
	"testing"

	"github.com/AlexeyNilov/gorpg/system/resource"
	"github.com/AlexeyNilov/gorpg/system/transformer"
	"github.com/stretchr/testify/assert"
)

func TestCreateLightEater(t *testing.T) {
	got := CreateLightEater()
	want := LightEater{
		transformer.ResourceTransformer{
			Resource: resource.Resource{
				Value:    0,
				MaxValue: 10,
			},
		},
	}
	assert.Equal(t, want, got)
}
