// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/4/24 下午7:33

package prototype

import (
	"testing"
)

func TestClone(t *testing.T) {
	shirtsCache := GetShirtsCloner()
	if shirtsCache == nil {
		t.Fatal("shirtCache was nil")
	}

	item1, err := shirtsCache.GetClone(White)
	if err != nil {
		t.Fatal(err)
	}

	if item1 == whitePrototype {
		t.Error("item1 cannot be equal to the whitePrototype")
	}

	shirt1, ok := item1.(*Shirt)
	if !ok {
		t.Fatal("Type assertion for shirt1 couldn't be done successfully")
	}

	shirt1.Price = 123.32
}