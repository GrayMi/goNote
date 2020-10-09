package main

import (
	"fmt"
)

type ST struct {
	AA int
	BB int
	CC int
	DD int
	EE int
}

func Modify9(a ST) ST {
	a.AA = 300
	return a
}

func Modify10(a ST) ST {
	b := a
	b.AA = 100
	b.BB = 500
	b.CC = 300
	return b
}

func Modify11(a ST) ST {
	b := new(ST)
	b = &a
	b.AA = 300
	return *b
}

func Modify12(a ST) ST {
	b := ST{}
	b = a
	b.AA = 300
	return b
}

func main() {
	a := ST{
		AA: 1,
		BB: 2,
		CC: 3,
		DD: 4,
		EE: 5,
	}
	b := Modify9(a)
	fmt.Println("a:", a) // 打印结果:{1 2 3 4 5}
	fmt.Println("b:", b) // 打印结果:{300 2 3 4 5}
	fmt.Println("")

	a1 := ST{
		AA: 1,
		BB: 2,
		CC: 3,
		DD: 4,
		EE: 5,
	}
	b1 := Modify10(a1)
	fmt.Println("a1:", a1) // 打印结果:{1 2 3 4 5}
	fmt.Println("b1:", b1) // 打印结果:{100 500 300 4 5}
	fmt.Println("")

	a2 := ST{
		AA: 1,
		BB: 2,
		CC: 3,
		DD: 4,
		EE: 5,
	}
	b2 := Modify11(a2)
	fmt.Println("a2:", a2) // 打印结果:{1 2 3 4 5}
	fmt.Println("b2:", b2) // 打印结果:{300 2 3 4 5}
	fmt.Println("")

	a3 := ST{
		AA: 1,
		BB: 2,
		CC: 3,
		DD: 4,
		EE: 5,
	}
	b3 := Modify12(a3)
	fmt.Println("a3:", a3) // 打印结果:{1 2 3 4 5}
	fmt.Println("b3:", b3) // 打印结果:{300 2 3 4 5}
	fmt.Println("")

	a4 := a3
	a4.AA = 100
	fmt.Println("a4:", a4) // 打印结果:{100 2 3 4 5}
}

/*
a: {1 2 3 4 5}
b: {300 2 3 4 5}

a1: {1 2 3 4 5}
b1: {100 500 300 4 5}

a2: {1 2 3 4 5}
b2: {300 2 3 4 5}

a3: {1 2 3 4 5}
b3: {300 2 3 4 5}

a4: {100 2 3 4 5}
*/

// Modify9 函数，直接操作，不会影响原参数
// Modify10函数，取值操作，不会影响原参数
// Modify11函数，取址操作，不会影响原参数
// Modify12函数，直接赋值，不会影响原参数
//
// 结论:以上结果说明 struct 的函数调用是传值
