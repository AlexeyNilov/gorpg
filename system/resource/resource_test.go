package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
    testResource := Resource{Value: 5}

    // Test getting 1 resource
    assert.Equal(t, 1, testResource.Get(1))
    assert.Equal(t, 4, testResource.Value)

    // Test getting the remaining 4 resources
    assert.Equal(t, 4, testResource.Get(5))
    assert.Equal(t, 0, testResource.Value)

    // Test getting negative value (should return 0)
    assert.Equal(t, 0, testResource.Get(-1))
}


func TestPut(t *testing.T) {
    testResource := Resource{Value: 5, MaxValue: 20}

    // Test putting 1 resource
    got := testResource.Put(1)
    assert.Equal(t, 1, got)
    assert.Equal(t, 6, testResource.Value)

    // Test putting 5 resources
    testResource.Put(5)
    assert.Equal(t, 11, testResource.Value)

    // Test putting negative amount (should not change value)
    got = testResource.Put(-1)
    assert.Equal(t, 0, got)
    assert.Equal(t, 11, testResource.Value)

    // Test putting more than max capacity
    got = testResource.Put(100)
    assert.Equal(t, 9, got)
    assert.Equal(t, 20, testResource.Value)  // Max value should be capped
}


