package structs

type CartInterface interface {
	AddProduct(product string)
	DeleteProduct(product string)
	ShowProduct()
}

type Cart struct {
	Name string
	Qty  int64
}

func (c Cart) AddProduct(product string) {

}

func (c Cart) DeleteProduct(product string) {

}

func (c Cart) ShowProduct() []*Cart {
	return []*Cart{}
}
