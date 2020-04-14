// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/4/9 上午10:17

package builder

// Product 一个具体的产品对象
type VehicleProduct struct {
	Wheels	int
	Seats 	int
	Structure 	string
}

// Builder 抽象建造者
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

// Director 指挥者
type VehicleBuildDirector struct {
	builder BuildProcess
}

// 指定在VehicleBuildDirector所使用的构造器
func (v *VehicleBuildDirector) SetBuilder(b BuildProcess) {
	v.builder = b
}

// 调用存储在VehicleBuildDirector的构造器，重现所有步骤
func (v *VehicleBuildDirector) Construct() {
	v.builder.SetWheels().SetSeats().SetStructure()
}

// ConCreateBuilder 具体建造者 a Builder of type Car
type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 4
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

// ConCreateBuilder 具体建造者 a Builder of type motorbike

type MotorbikeBuilder struct{
	v VehicleProduct
}

func (m *MotorbikeBuilder) SetWheels() BuildProcess {
	m.v.Wheels = 2
	return m
}

func (m *MotorbikeBuilder) SetSeats() BuildProcess {
	m.v.Seats = 2
	return m
}

func (m *MotorbikeBuilder) SetStructure() BuildProcess {
	m.v.Structure = "MotorBike"
	return m
}

func (m *MotorbikeBuilder) GetVehicle() VehicleProduct {
	return m.v
}



