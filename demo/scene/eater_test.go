package main

import (
	"testing"

	"github.com/AlexeyNilov/gorpg/system/resource"
	"github.com/AlexeyNilov/gorpg/system/transformer"
	"github.com/stretchr/testify/assert"
)

func TestNewLightEater(t *testing.T) {
	got := NewLightEater()
	want := LightEater{
		ResourceTransformer: transformer.ResourceTransformer{
			Resource: resource.Resource{
				Value:    0,
				MaxValue: 10,
			},
		},
		ConsumeRate: 2,
		ProduceRate: 1,
	}

	assert.Equal(t, want, got)
}

func TestExecute(t *testing.T) {
	eater := NewLightEater()
	resource := resource.Resource{
		Value:    5,
		MaxValue: 10,
	}
	eater.Execute("Light", &resource, &resource)

	assert.Equal(t, 4, resource.Value)
	assert.Equal(t, 1, eater.Resource.Value)
	assert.Equal(t, 1, eater.ConsumedCount)
	assert.Equal(t, 1, eater.ProducedCount)

	eater = NewLightEater()
	eater.Execute("Night", &resource, &resource)
	assert.Equal(t, 4, resource.Value)
	assert.Equal(t, 0, eater.Resource.Value)
	assert.Equal(t, 0, eater.ConsumedCount)
	assert.Equal(t, 0, eater.ProducedCount)
}

func BenchmarkExecute(b *testing.B) {
	resource := resource.Resource{
		Value:    100,
		MaxValue: 100,
	}
	eaters := []LightEater{}
	for range 100 {
		eaters = append(eaters, NewLightEater())
	}

	b.ResetTimer()
	for i := range eaters {
		eaters[i].Execute("Light", &resource, &resource)
	}
}
