package div_test

import (
	"go-tests/pkg/div"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivisao(t *testing.T) {
	result, err := div.Divide(6, 2)
	expect := 3

	assert.Equal(t, expect, result, "o resultado atual precisa ser igual ao resultado esperado")
	assert.Nil(t, err)
}

func TestDivisao2(t *testing.T) {
	result, err := div.Divide(6, 0)
	expect := "O denominador não pode ser 0"

	assert.Equal(t, 0, result, "o resultado atual precisa ser igual ao resultado esperado")
	assert.Equal(t, expect, err.Error())
	assert.NotNil(t, err)
}
