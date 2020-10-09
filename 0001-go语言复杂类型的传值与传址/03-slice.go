package main

import (
	"fmt"
)

func Modify9(a []int) []int {
	a[1] = 100
	b := append(a[:0:0], a...)
	b[0] = 300
	return b
}

func Modify10(a []int) []int {
	c := a
	c[0] = 100
	b := append(c, 500)
	b[1] = 300
	return b
}

// 为了防止slice的魔幻操作，如果运行服务器不是太差的，请先复制操作
// 然后统一当做传值处理
func Modify11(a []int) []int {
	c := make([]int, len(a))
	c = a
	c[0] = 100
	b := append(c, 500)
	b[1] = 300
	return b
}

func Modify12(a []int) []int {
	c := make([]int, len(a))
	copy(c, a)
	c[0] = 100
	b := append(c, 500)
	b[1] = 300
	return b
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := Modify9(a)
	fmt.Println("a:", a) // 打印结果:[1 2 3 4 5]
	fmt.Println("b:", b) // 打印结果:[300 2 3 4 5]
	fmt.Println("")

	a1 := []int{1, 2, 3, 4, 5}
	b1 := Modify10(a1)
	fmt.Println("a1:", a1) // 打印结果:[100 2 3 4 5]
	fmt.Println("b1:", b1) // 打印结果:[100 300 3 4 5 500]
	fmt.Println("")

	a2 := []int{1, 2, 3, 4, 5}
	b2 := Modify11(a2)
	fmt.Println("a2:", a2) // 打印结果:[100 2 3 4 5]
	fmt.Println("b2:", b2) // 打印结果:[100 300 3 4 5 500]
	fmt.Println("")

	a3 := []int{1, 2, 3, 4, 5}
	b3 := Modify12(a3)
	fmt.Println("a3:", a3) // 打印结果:[1 2 3 4 5]
	fmt.Println("b3:", b3) // 打印结果:[100 300 3 4 5 500]
	fmt.Println("")

}

/*
a: [1 100 3 4 5]
b: [300 2 3 4 5]

a1: [100 2 3 4 5]
b1: [100 300 3 4 5 500]

a2: [100 2 3 4 5]
b2: [100 300 3 4 5 500]

a3: [1 2 3 4 5]
b3: [100 300 3 4 5 500]
*/

// Modify9 函数，深拷贝的第二种方式 - 触发值赋值，  不会影响原参数
// Modify10函数，取值操作，  会影响原参数
// Modify11函数，取址操作，  会影响原参数
// Modify12函数，copy操作后，不会影响原参数

// 结论:以上结果说明 slice 的函数调用是传址，但在append扩容后，新slice不会影响原参数
