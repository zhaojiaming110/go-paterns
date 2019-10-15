package car

type Director struct {
	builder Builder
}

func (d *Director) SetBuilder(builder Builder) {
	d.builder = builder
}

func (d *Director) Create(color, wheels string, speed float64) *Car {
	d.builder.BuildColor(color)
	d.builder.BuildSpeed(speed)
	d.builder.BuildWheels(wheels)
	return d.builder.GetResult().(*Car)
}
