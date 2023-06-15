package calc_test

import (
	"go-tests/pkg/calc"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSub(t *testing.T) {
	result := calc.Sub(5, 3)
	expect := 2

	assert.Equal(t, expect, result, "o resultado atual precisa ser igual ao resultado esperado")
}

func TestSub2(t *testing.T) {
	result := calc.Sub(5, 3)
	expect := 0

	assert.Equal(t, expect, result, "o resultado atual precisa ser igual ao resultado esperado")
}
