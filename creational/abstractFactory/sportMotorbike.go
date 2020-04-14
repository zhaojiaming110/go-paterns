// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/4/14 上午10:50

package abstractFactory

type SportMotorbike struct{}

func (s *SportMotorbike) GetWheels() int {
	return 2
}
func (s *SportMotorbike) GetSeats() int {
	return 1
}
func (s *SportMotorbike) GetType() int {
	return SportMotorbikeType
}