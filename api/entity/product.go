package entity

type Product struct {
	Id    uint32
	Name  string
	Count uint32
}

func NewProduct(name string, count uint32) *Product {
	return &Product{
		Name:  name,
		Count: count,
	}
}
