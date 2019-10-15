package kfc

import "fmt"

type Kfc struct {
	Food  string
	Drink string
}

func (k *Kfc) Show() {
	fmt.Println("本套餐为:"+k.Food+"+", k.Drink)
}
