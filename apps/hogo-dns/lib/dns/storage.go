package dns

type StorageAccessor interface {
	Get() interface{}
}

type StorageWriter interface {
	Store()
}

type Storage interface {
	StorageAccessor
	StorageWriter
}
