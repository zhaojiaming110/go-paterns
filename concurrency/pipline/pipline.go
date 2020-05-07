// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/7 下午3:14

package main

import "fmt"

func generator(max int) <-chan int{
	out := make(chan int, 100)
	go func() {
		for i := 1; i <= max; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func power(in <-chan int) <-chan int{
	out := make(chan int, 100)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func sum(in <-chan int) <-chan int{
	out := make(chan int, 100)
	go func() {
		var sum int
		for v := range in {
			sum += v
		}
		out <- sum
		close(out)
	}()
	return out
}

func main() {
	// [1, 2, 3]
	fmt.Println(<-sum(power(generator(3))))
}
