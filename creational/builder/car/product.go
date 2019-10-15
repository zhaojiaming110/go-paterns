package car

import "fmt"

type Car struct {
	Speed  float64
	Color  string
	Wheels string
}

func (c *Car) Show() {
	fmt.Println("this car's Color:"+c.Color+" wheels:"+c.Wheels+" speed:", c.Speed)
}
