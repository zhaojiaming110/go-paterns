// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/4/9 下午3:26

package factoryMethod

import (
	"testing"
)

func TestGetPaymentMethodCash(t *testing.T) {
	paymet, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("A Payment method of type 'Cash' must exit")
	}
	msg := paymet.Pay(11.22)
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodDebitCard(t *testing.T) {
	paymet, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Fatal("A Payment method of type 'Cash' must exit")
	}
	msg := paymet.Pay(33.18)
	t.Log("LOG:", msg)
}