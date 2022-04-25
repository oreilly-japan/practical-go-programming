package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Add(a, b int) (int, error) {
	return a + b, nil
}

// testify-test
func TestByTestify(t *testing.T) {
	result, err := Add(1, 2)
	assert.Nil(t, err)
	assert.Equal(t, 3, result)
}

// testify-test

func TestErrorMessage(t *testing.T) {
	assert.Equal(t, 2, 3)
	type Person struct {
		Name string
		Age  int
	}
	assert.Equal(t, Person{"織田信長", 49}, Person{"徳川家康", 73})
}
