package transformer

import (
	"testing"

	"github.com/AlexeyNilov/gorpg/stuff/resource"
	"github.com/stretchr/testify/assert"
)

func TestConsume(t *testing.T) {
	testTransformer := Transformer{}
	testTransformer.Consume()
	assert.Equal(t, 1, testTransformer.Consumed)
}

func TestProduce(t *testing.T) {
	testTransformer := Transformer{}
	testTransformer.Produce()
	assert.Equal(t, 1, testTransformer.Produced)
}

func TestPut(t *testing.T) {
	testResource := resource.Resource{Value: 5, MaxValue: 20}
	testResourceTransformer := ResourceTransformer{}
	testResourceTransformer.Consume(&testResource, 5)

	assert.Equal(t, 0, testResource.Value)
	assert.Equal(t, 1, testResourceTransformer.ConsumedCount)
	assert.Equal(t, 5, testResourceTransformer.Capacity)

	testResourceTransformer.Produce(&testResource, 5)

	assert.Equal(t, 5, testResource.Value)
	assert.Equal(t, 1, testResourceTransformer.ProducedCount)
	assert.Equal(t, 0, testResourceTransformer.Capacity)
}
