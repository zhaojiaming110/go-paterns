// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/4/14 上午10:40

package abstractFactory

type FamiliarCar struct{}

func (f *FamiliarCar) GetWheels() int {
	return 4
}

func (f *FamiliarCar) GetSeats() int {
	return 5
}

func (f *FamiliarCar) GetDoors() int {
	return 4
}