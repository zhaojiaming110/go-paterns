// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/4/9 上午11:07

package builder

import (
	"testing"
)

func TestBuilderPattern(t *testing.T) {
	// 获取director实例
	vehicleBuildDirector := GetvehicleBuildDirector()

	carBuilder := &CarBuilder{}
	vehicleBuildDirector.SetBuilder(carBuilder)
	vehicleBuildDirector.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Error("wheels must be 4")
	}

	if car.Seats != 4 {
		t.Error("seats must be 4")
	}

	if car.Structure != "Car" {
		t.Error("strutures must be Car")
	}
}