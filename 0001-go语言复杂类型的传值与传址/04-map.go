package main

import (
	"fmt"
)

func Modify9(a map[string]int32) map[string]int32 {
	a["0"] = 300
	return a
}

func Modify10(a map[string]int32) map[string]int32 {
	b := a
	b["0"] = 100
	b["1"] = 300
	b["5"] = 500
	return b
}

func Modify11(a map[string]int32) map[string]int32 {
	b := make(map[string]int32, len(a))
	b = a
	b["0"] = 300
	return b
}

func Modify12(a map[string]int32) map[string]int32 {
	b := map[string]int32{}
	for k, v := range a {
		b[k] = v
	}

	b["0"] = 300
	return b
}

func main() {
	a := map[string]int32{
		"0": 1,
		"1": 2,
		"2": 3,
		"3": 4,
		"4": 5,
	}
	b := Modify9(a)
	fmt.Println("a:", a) // 打印结果:map[0:300 1:2 2:3 3:4 4:5]
	fmt.Println("b:", b) // 打印结果:map[0:300 1:2 2:3 3:4 4:5]
	fmt.Println("")

	a1 := map[string]int32{
		"0": 1,
		"1": 2,
		"2": 3,
		"3": 4,
		"4": 5,
	}
	b1 := Modify10(a1)
	fmt.Println("a1:", a1) // 打印结果:map[0:100 1:300 2:3 3:4 4:5 5:500]
	fmt.Println("b1:", b1) // 打印结果:map[0:100 1:300 2:3 3:4 4:5 5:500]
	fmt.Println("")

	a2 := map[string]int32{
		"0": 1,
		"1": 2,
		"2": 3,
		"3": 4,
		"4": 5,
	}
	b2 := Modify11(a2)
	fmt.Println("a2:", a2) // 打印结果:map[0:300 1:2 2:3 3:4 4:5]
	fmt.Println("b2:", b2) // 打印结果:map[0:300 1:2 2:3 3:4 4:5]
	fmt.Println("")

	a3 := map[string]int32{
		"0": 1,
		"1": 2,
		"2": 3,
		"3": 4,
		"4": 5,
	}
	b3 := Modify12(a3)
	fmt.Println("a3:", a3) // 打印结果:map[0:1 1:2 2:3 3:4 4:5]
	fmt.Println("b3:", b3) // 打印结果:map[0:300 1:2 2:3 3:4 4:5]
	fmt.Println("")

}

/*
a: map[0:300 1:2 2:3 3:4 4:5]
b: map[0:300 1:2 2:3 3:4 4:5]

a1: map[0:100 1:300 2:3 3:4 4:5 5:500]
b1: map[0:100 1:300 2:3 3:4 4:5 5:500]

a2: map[0:300 1:2 2:3 3:4 4:5]
b2: map[0:300 1:2 2:3 3:4 4:5]

a3: map[0:1 1:2 2:3 3:4 4:5]
b3: map[0:300 1:2 2:3 3:4 4:5]

*/
// Modify9 函数，直接操作，  会影响原参数
// Modify10函数，取值操作，  会影响原参数
// Modify11函数，取址操作，  会影响原参数
// Modify12函数，深度复制后，不会影响原参数
//
// 结论:以上结果说明 map 的函数调用是传址，除非深度复制后，基本上一改所有的都跟着改
