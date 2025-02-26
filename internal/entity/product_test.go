package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct_New(t *testing.T) {
	p, err := NewProduct("Computer", 1000.00)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "Computer", p.Name)
	assert.Equal(t, 1000.00, p.Price)
}

func TestProduct_Validate(t *testing.T) {
	p, err := NewProduct("Computer", 1000.00)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 1000.00)
	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Computer", 0)
	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("Computer", -10)
	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsInvalid, err)
}
