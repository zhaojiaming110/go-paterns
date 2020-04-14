// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/4/9 上午11:29

// 将Builder中的director设计为Singleton模式
package builder

type vehicleBuildDirector struct {
	builder BuildProcess
}

var instance *vehicleBuildDirector

// 对外提供vehicleBuildDirector唯一的实例
func GetvehicleBuildDirector() *vehicleBuildDirector {
	if instance == nil {
		instance = new(vehicleBuildDirector)
	}

	return instance
}

// 指定在VehicleBuildDirector所使用的构造器
func (v *vehicleBuildDirector) SetBuilder(b BuildProcess) {
	v.builder = b
}

// 调用存储在VehicleBuildDirector的构造器，重现所有步骤
func (v *vehicleBuildDirector) Construct() {
	v.builder.SetWheels().SetSeats().SetStructure()
}