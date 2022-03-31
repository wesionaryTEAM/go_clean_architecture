package sample_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	got := math.Abs(-1)
	assert.Equal(t, got, float64(1), "they should be equal")
}
