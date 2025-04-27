package transformer

import (
	"testing"

	"github.com/AlexeyNilov/gorpg/system/resource"
	"github.com/stretchr/testify/assert"
)

func TestConsume(t *testing.T) {
	testTransformer := Transformer{}
	testTransformer.Consume()
	assert.Equal(t, 1, testTransformer.ConsumedCount)
}

func TestProduce(t *testing.T) {
	testTransformer := Transformer{}
	testTransformer.Produce()
	assert.Equal(t, 1, testTransformer.ProducedCount)
}

func TestResourceConsume(t *testing.T) {
	// Initialize resources and transformer
	testResource := resource.Resource{Value: 5, MaxValue: 20}
	testTransformer := ResourceTransformer{
		Transformer: Transformer{},
		Resource:    resource.Resource{Value: 0, MaxValue: 10},
	}

	// Test consuming resources
	testTransformer.Consume(&testResource, 5)
	assert.Equal(t, 0, testResource.Value, "Resource value should be depleted after consuming 5")
	assert.Equal(t, 1, testTransformer.ConsumedCount, "Transformer should record one consumption")
	assert.Equal(t, 5, testTransformer.Value, "Transformer value should increase by consumed amount")

	// Test buffer overflow during consumption
	testResource = resource.Resource{Value: 15, MaxValue: 20}
	testTransformer.Value = 0
	testTransformer.Consume(&testResource, 15)
	assert.Equal(t, 10, testTransformer.Value, "Transformer should cap value at max capacity (10)")
	assert.Equal(t, 0, testResource.Value, "Excess resource should be wasted")

}

func TestResourceProduce(t *testing.T) {
	// Initialize resources and transformer
	testResource := resource.Resource{Value: 0, MaxValue: 20}
	testTransformer := ResourceTransformer{
		Transformer: Transformer{},
		Resource:    resource.Resource{Value: 5, MaxValue: 10},
	}

	// Test producing resources
	testTransformer.Produce(&testResource, 5)
	assert.Equal(t, 5, testResource.Value, "Resource value should increase by 5 after production")
	assert.Equal(t, 1, testTransformer.ProducedCount, "Transformer should record one production")
	assert.Equal(t, 0, testTransformer.Value, "Transformer value should decrease after production")

	// Test buffer overflow during production
	testResource = resource.Resource{Value: 0, MaxValue: 20}
	testTransformer.Value = 10
	testTransformer.Produce(&testResource, 15)
	assert.Equal(t, 0, testTransformer.Value, "Transformer value should be depleted after production")
	assert.Equal(t, 10, testResource.Value, "Resource value should cap at max Transformer capacity (10)")
}
