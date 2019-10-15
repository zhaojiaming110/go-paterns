package car

type Builder interface {
	BuildSpeed(speed float64) Builder
	BuildColor(color string) Builder
	BuildWheels(wheels string) Builder
	GetResult() interface{}
}
