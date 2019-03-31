package storage

type StorageMock struct{}

func NewStorageMock() Storage {
	return &StorageMock{}
}

func (storage *StorageMock) Put(key string, file []byte) error {
	return nil
}

func (storage *StorageMock) Destroy() error {
	return nil
}
