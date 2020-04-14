// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/4/14 上午10:43

package abstractFactory

type LuxuryCar struct{}

func (l *LuxuryCar) GetWheels() int {
	return 4
}

func (l *LuxuryCar) GetSeats() int {
	return 5
}

func (l *LuxuryCar) GetDoors() int {
	return 4
}


