package kfc

type Builder interface {
	NewProduct()
	BuildFood() Builder
	BuildDrink() Builder
	GetResult() interface{}
}
