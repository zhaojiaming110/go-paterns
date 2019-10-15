package kfc

// 套餐A
type KfcBuilderA struct {
	Kfc *Kfc
}

func (kb *KfcBuilderA) NewProduct() {
	kb.Kfc = new(Kfc)
}

func (kb *KfcBuilderA) BuildFood() Builder {
	kb.Kfc.Food = "牛肉汉堡"
	return kb
}

func (kb *KfcBuilderA) BuildDrink() Builder {
	kb.Kfc.Drink = "柠檬汁"
	return kb
}

func (kb *KfcBuilderA) GetResult() interface{} {
	return kb.Kfc
}

// 套餐B
type KfcBuilderB struct {
	Kfc *Kfc
}

func (kb *KfcBuilderB) NewProduct() {
	kb.Kfc = new(Kfc)
}

func (kb *KfcBuilderB) BuildFood() Builder {
	kb.Kfc.Food = "鸡肉卷"
	return kb
}

func (kb *KfcBuilderB) BuildDrink() Builder {
	kb.Kfc.Drink = "酸奶"
	return kb
}

func (kb *KfcBuilderB) GetResult() interface{} {
	return kb.Kfc
}
