// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/4/9 下午2:59

package factoryMethod

import (
	"errors"
	"fmt"
)

// PaymentMethod 定义了在商店中付款的支付方式
// 此Factory Method返回实现此接口的对象
type PaymentMethod interface {
	Pay(amount float32) string
}

// 我们目前实现的支付方式如下
const (
	Cash = 1
	DebitCard = 2
)

// GetPaymentMethod returns a pointer to a PaymentMethod object or an error
func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized\n", m))
	}
}

type CashPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f payed using cash\n", amount)
}

type DebitCardPM struct{}

func (d *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f payed using debit card\n", amount)
}

