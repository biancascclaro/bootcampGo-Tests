package tests

import (
	"go-tests-aula2-morning/internal/products"
	"go-tests-aula2-morning/tests/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateService(t *testing.T) {
	storeMock := mock.StorageMock{}
	repository := products.NewRepository(&storeMock)
	service := products.NewService(repository)

	expected := products.Product{
		Id:    1,
		Name:  "After Update",
		Price: 20,
	}

	product, err := service.Update(1, "After Update", 20)

	assert.Equal(t, expected, product)
	assert.True(t, storeMock.ReadCalled)
	assert.Nil(t, err)
}

func TestDeleteServiceProductExists(t *testing.T) {
	storeMock := mock.StorageMock{}
	repository := products.NewRepository(&storeMock)
	service := products.NewService(repository)

	err := service.Delete(1)

	assert.NoError(t, err)
	assert.True(t, storeMock.ReadCalled)
	assert.False(t, storeMock.WriteCalled)
}

func TestDeleteServiceProductNotExists(t *testing.T) {
	storeMock := mock.StorageMock{}
	repository := products.NewRepository(&storeMock)
	service := products.NewService(repository)

	err := service.Delete(2)

	assert.Error(t, err)
	assert.True(t, storeMock.ReadCalled)
	assert.False(t, storeMock.WriteCalled)
}
