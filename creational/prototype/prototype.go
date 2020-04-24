// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/4/24 下午7:09

package prototype

import (
	"errors"
	"fmt"
)

// 定义一个衬衫克隆的接口
type ShirtCloner interface {
	GetClone(m int) (ItemInfoGetter, error)
}

// 定义一个包级函数用来获取实现ShirtCloner的具体函数
func GetShirtsCloner() ShirtCloner {
	return new(ShirtsCache)
}

// 实现ShirtCloner接口
type ShirtsCache struct{}

const (
	White = 1
	Black = 2
	Blue = 3
)

func (s *ShirtsCache) GetClone(m int) (ItemInfoGetter, error) {
	switch m {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("Shirt model not recognized")
	}
}

type ItemInfoGetter interface {
	GetInfo() string
}

type Shirt struct{
	Price	float32
	Color 	int
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt Price is %f, Color is %d\n", s.Price, s.Color)
}

var whitePrototype *Shirt = &Shirt{
	Price: 100,
	Color: White,
}

var blackPrototype *Shirt = &Shirt{
	Price: 70,
	Color: Black,
}

var bluePrototype *Shirt = &Shirt{
	Price: 80,
	Color: Blue,
}










