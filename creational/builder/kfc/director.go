package kfc

type Director struct {
	builder Builder
}

func (d *Director) SetBuilder(builder Builder) {
	d.builder = builder
}

func (d *Director) Create() *Kfc {
	d.builder.NewProduct()
	d.builder.BuildFood()
	d.builder.BuildDrink()
	return d.builder.GetResult().(*Kfc)
}
