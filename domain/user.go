package domain

type UserWithInterface struct {
	ID   uint64
	Name string
	Data interface{}
}

type CustomData interface {
	int64 | DataStructure
}

type DataStructure struct {
	Age    uint
	Height float64
}

type UserGeneric[T CustomData] struct {
	ID   uint64
	Name string
	Data T
}
