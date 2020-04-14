// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/4/14 上午10:51

package abstractFactory

type CruiseMotorbike struct{}

func (c *CruiseMotorbike) GetWheels() int {
	return 2
}

func (c *CruiseMotorbike) GetSeats() int {
	return 2
}

func (c *CruiseMotorbike) GetType() int {
	return CruiseMotorbikeType
}