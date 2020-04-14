// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/4/14 上午10:30

package abstractFactory

import (
	"errors"
	"fmt"
)

const (
	LuxuryCarType   = 1
	FamiliarCarType = 2
)

type CarFactory struct{}

func (c *CarFactory) GetVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(FamiliarCar), nil
	case FamiliarCarType:
		return new(LuxuryCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized\n", v))
	}
}
