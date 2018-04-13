// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"
)

// fibonacci is a function that returns
// a function that returns an int.
// 结论， 这个不是一个递归的fibonacci
func fibonacci() func() int {
	f, g := 1, 0 // 这个是一个内部的全局变量
	fmt.Printf("%d, %d\n", f, g)
	return func() int {
		fmt.Printf("= %d, %d", f, g)
		f, g = g, f+g // 这个是一个交换算法， 可以用于排序
		fmt.Printf(", %d + %d = %d\n", f, g, f+g)
		return f
	}
}

func main() {
	f := fibonacci() // 这个函数是一个指针 (关键是这个)
	for i := 0; i < 10; i++ {
		fmt.Printf("i %d: ", i)
		//fmt.Println(fibonacci())
		fmt.Println(f())
	}
}
