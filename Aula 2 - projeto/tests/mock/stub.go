package mock

import "encoding/json"

type Product struct {
	Id    int
	Name  string
	Price float64
}

type StorageStub struct{}

func (s *StorageStub) Read(data interface{}) error {
	products := []Product{
		{Id: 1, Name: "café", Price: 10},
		{Id: 2, Name: "suco", Price: 15},
	}

	encoded, _ := json.Marshal(products)
	return json.Unmarshal(encoded, &data)
}

func (s *StorageStub) Write(data interface{}) error {
	return nil
}
