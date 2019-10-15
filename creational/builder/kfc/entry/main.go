package main

import "go-paterns/creational/builder/kfc"

func main() {
	builderA := new(kfc.KfcBuilderA)
	director := new(kfc.Director)
	director.SetBuilder(builderA)
	kfcmeal := director.Create()
	kfcmeal.Show()

	builderB := new(kfc.KfcBuilderB)
	director.SetBuilder(builderB)
	kfcmeal = director.Create()
	kfcmeal.Show()
}
