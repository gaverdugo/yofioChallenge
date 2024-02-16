package creditassigner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssignNoValidCombination(t *testing.T) {
	assert := assert.New(t)

	var target int32 = 400
	creditAssigner := CreditAssigner{}

	a1, a2, a3, err := creditAssigner.Assign(target)

	var ex1, ex2, ex3 int32 = 0, 0, 0

	assert.Equal(ex1, a1, "first value should be 0")
	assert.Equal(ex2, a2, "second value should be 0")
	assert.Equal(ex3, a3, "third value should be 0")
	assert.NotNil(err, "function should throw an error")
}

func TestAssignValidCombination(t *testing.T) {
	assert := assert.New(t)

	var target int32 = 6700
	creditAssigner := CreditAssigner{}

	a1, a2, a3, err := creditAssigner.Assign(target)

	var ex1, ex2, ex3 int32 = 2, 1, 8

	assert.Equal(ex1, a1, "should assign 2 300 credits")
	assert.Equal(ex2, a2, "should assign 1 500 credit")
	assert.Equal(ex3, a3, "should assign 8 700 credits")
	assert.Nil(err, "function should not throw an error")
}
