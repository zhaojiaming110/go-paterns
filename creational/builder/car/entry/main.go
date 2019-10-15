package main

import (
	"go-paterns/creational/builder/car"
)

func main() {
	builder := new(car.CarBuilder)
	director := new(car.Director)
	director.SetBuilder(builder)
	car := director.Create("red", "sport", 180.64)
	car.Show()
}
