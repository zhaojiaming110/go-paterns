package car

type CarBuilder struct {
	Car *Car
}

func (cb *CarBuilder) BuildSpeed(speed float64) Builder {

	if cb.Car == nil {
		cb.Car = new(Car)
	}
	cb.Car.Speed = speed

	return cb
}

func (cb *CarBuilder) BuildColor(color string) Builder {

	if cb.Car == nil {
		cb.Car = new(Car)
	}

	cb.Car.Color = color

	return cb
}

func (cb *CarBuilder) BuildWheels(wheels string) Builder {
	if cb.Car == nil {
		cb.Car = new(Car)
	}

	cb.Car.Wheels = wheels

	return cb
}

func (cb *CarBuilder) GetResult() interface{} {
	return cb.Car
}
