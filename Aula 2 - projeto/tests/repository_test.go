package tests

import (
	"go-tests-aula2-morning/internal/products"
	"go-tests-aula2-morning/tests/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StoreMock struct{}

func TestGetAllRepo(t *testing.T) {
	storeStub := mock.StorageStub{}
	repository := products.NewRepository(&storeStub)

	expected := []products.Product{
		{Id: 1, Name: "café", Price: 10},
		{Id: 2, Name: "suco", Price: 15},
	}

	products, err := repository.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, products, expected)
}

func TestUpdateRepo(t *testing.T) {
	storeMock := mock.StorageMock{}
	repository := products.NewRepository(&storeMock)

	expected := products.Product{
		Id:    1,
		Name:  "After Update",
		Price: 20,
	}

	product, err := repository.Update(1, "After Update", 20)

	assert.Equal(t, expected, product)
	assert.True(t, storeMock.ReadCalled)
	assert.Nil(t, err)
}
