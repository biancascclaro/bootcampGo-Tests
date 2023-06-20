package mock

import "encoding/json"

type StorageMock struct {
	ReadCalled  bool
	WriteCalled bool
}

func (s *StorageMock) Read(data interface{}) error {
	s.ReadCalled = true

	encoded, _ := json.Marshal([]Product{{
		Id:    1,
		Name:  "Before Update",
		Price: 5,
	}})

	return json.Unmarshal(encoded, &data)
}

func (s *StorageMock) Write(data interface{}) error {
	s.WriteCalled = true

	return nil
}
