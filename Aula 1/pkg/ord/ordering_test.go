package ord_test

import (
	"go-tests/pkg/ord"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrd(t *testing.T) {
	nums := []int{10, 500, 48, 0, -3, 2, 185}
	result := ord.Ord(nums)
	expect := []int{-3, 0, 2, 10, 48, 185, 500}

	assert.Equal(t, expect, result, "o resultado atual precisa ser igual ao resultado esperado")
}
