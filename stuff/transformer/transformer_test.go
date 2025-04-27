package transformer

import (
	"testing"

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

