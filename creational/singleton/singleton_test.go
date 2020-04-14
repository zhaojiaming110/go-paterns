// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/4/8 下午1:12

package singleton

import (
	"testing"
)

func Test_singleton_GetInstance(t *testing.T) {
	couter1 := GetInstance()
	if couter1 == nil {
		t.Error("A new connection object must have been made")
	}

	if couter1.AddOne() != 1 {
		t.Error("the count must be 1 ")
	}

	counter2 := GetInstance()
	if couter1 != counter2 {
		t.Error("Singleton instences must be same")
	}

	if counter2.AddOne() != 2 {
		t.Error("the count must be 2")
	}
}