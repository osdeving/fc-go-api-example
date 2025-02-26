package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct_New(t *testing.T) {
	product, _ := NewProduct("Computer", 1000.00)

	assert.NotNil(t, product)
}

