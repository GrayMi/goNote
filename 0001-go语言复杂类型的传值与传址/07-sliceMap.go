package main

import (
	"fmt"
)

func Modify9(a []map[string]int) []map[string]int {
	a[0]["AA"] = 300
	return a
}

func Modify10(a []map[string]int) map[string]int {
	b := a[0]
	b["AA"] = 100
	b["BB"] = 500
	b["CC"] = 300
	return b
}

func Modify11(a []map[string]int) map[string]int {
	b := make(map[string]int, 5)
	b = a[0]

	b["AA"] = 300
	return b
}

func Modify12(a []map[string]int) map[string]int {
	b := map[string]int{}
	b = a[0]
	b["AA"] = 300
	return b
}

func Modify13(a []map[string]int) []map[string]int {
	b := make([]map[string]int, 2)
	copy(b, a)
	b[0]["AA"] = 300
	return b
}

func Modify14(a []map[string]int) []map[string]int {
	d := map[string]int{
		"AA": 1,
		"BB": 2,
		"CC": 3,
		"DD": 4,
		"EE": 5,
	}
	b := append(a, d)
	b[0]["AA"] = 300
	return b
}

func Modify15(a []map[string]int) []map[string]int {
	d := map[string]int{
		"AA": 1,
		"BB": 2,
		"CC": 3,
		"DD": 4,
		"EE": 5,
	}
	b := append(a, d)
	aMap := map[string]int{}
	for k, v := range b[0] {
		aMap[k] = v
	}

	aMap["AA"] = 300

	b[0] = aMap
	return b
}

func main() {
	a := []map[string]int{
		{
			"AA": 1,
			"BB": 2,
			"CC": 3,
			"DD": 4,
			"EE": 5,
		},
		{
			"AA": 1,
			"BB": 2,
			"CC": 3,
			"DD": 4,
			"EE": 5,
		},
	}
	b := Modify9(a)
	fmt.Println("a:", a) // 打印结果:[map[AA:300 BB:2 CC:3 DD:4 EE:5] map[AA:1 BB:2 CC:3 DD:4 EE:5]]
	fmt.Println("b:", b) // 打印结果:[map[AA:300 BB:2 CC:3 DD:4 EE:5] map[AA:1 BB:2 CC:3 DD:4 EE:5]]
	fmt.Println("")

	a1 := []map[string]int{
		{
			"AA": 1,
			"BB": 2,
			"CC": 3,
			"DD": 4,
			"EE": 5,
		},
		{
			"AA": 1,
			"BB": 2,
			"CC": 3,
			"DD": 4,
			"EE": 5,
		},
	}
	b1 := Modify10(a1)
	fmt.Println("a1:", a1) // 打印结果:[map[AA:100 BB:500 CC:300 DD:4 EE:5] map[AA:1 BB:2 CC:3 DD:4 EE:5]]
	fmt.Println("b1:", b1) // 打印结果:map[AA:100 BB:500 CC:300 DD:4 EE:5]
	fmt.Println("")

	a2 := []map[string]int{
		{
			"AA": 1,
			"BB": 2,
			"CC": 3,
			"DD": 4,
			"EE": 5,
		},
		{
			"AA": 1,
			"BB": 2,
			"CC": 3,
			"DD": 4,
			"EE": 5,
		},
	}
	b2 := Modify11(a2)
	fmt.Println("a2:", a2) // 打印结果:[map[AA:300 BB:2 CC:3 DD:4 EE:5] map[AA:1 BB:2 CC:3 DD:4 EE:5]]
	fmt.Println("b2:", b2) // 打印结果:map[AA:300 BB:2 CC:3 DD:4 EE:5]
	fmt.Println("")

	a3 := []map[string]int{
		{
			"AA": 1,
			"BB": 2,
			"CC": 3,
			"DD": 4,
			"EE": 5,
		},
		{
			"AA": 1,
			"BB": 2,
			"CC": 3,
			"DD": 4,
			"EE": 5,
		},
	}
	b3 := Modify12(a3)
	fmt.Println("a3:", a3) // 打印结果: [map[AA:300 BB:2 CC:3 DD:4 EE:5] map[AA:1 BB:2 CC:3 DD:4 EE:5]]
	fmt.Println("b3:", b3) // 打印结果: map[AA:300 BB:2 CC:3 DD:4 EE:5]
	fmt.Println("")

	a4 := []map[string]int{
		{
			"AA": 1,
			"BB": 2,
			"CC": 3,
			"DD": 4,
			"EE": 5,
		},
		{
			"AA": 1,
			"BB": 2,
			"CC": 3,
			"DD": 4,
			"EE": 5,
		},
	}
	b4 := Modify13(a4)
	fmt.Println("a4:", a4) // 打印结果: [map[AA:300 BB:2 CC:3 DD:4 EE:5] map[AA:1 BB:2 CC:3 DD:4 EE:5]]
	fmt.Println("b4:", b4) // 打印结果: [map[AA:300 BB:2 CC:3 DD:4 EE:5] map[AA:1 BB:2 CC:3 DD:4 EE:5]]
	fmt.Println("")

	a5 := []map[string]int{
		{
			"AA": 1,
			"BB": 2,
			"CC": 3,
			"DD": 4,
			"EE": 5,
		},
		{
			"AA": 1,
			"BB": 2,
			"CC": 3,
			"DD": 4,
			"EE": 5,
		},
	}
	b5 := Modify14(a5)
	fmt.Println("a5:", a5) // 打印结果: [map[AA:300 BB:2 CC:3 DD:4 EE:5] map[AA:1 BB:2 CC:3 DD:4 EE:5]]
	fmt.Println("b5:", b5) // 打印结果: [map[AA:300 BB:2 CC:3 DD:4 EE:5] map[AA:1 BB:2 CC:3 DD:4 EE:5] map[AA:1 BB:2 CC:3 DD:4 EE:5]]
	fmt.Println("")

	a6 := []map[string]int{
		{
			"AA": 1,
			"BB": 2,
			"CC": 3,
			"DD": 4,
			"EE": 5,
		},
		{
			"AA": 1,
			"BB": 2,
			"CC": 3,
			"DD": 4,
			"EE": 5,
		},
	}
	b6 := Modify15(a6)
	fmt.Println("a6:", a6) // 打印结果: [map[AA:1 BB:2 CC:3 DD:4 EE:5] map[AA:1 BB:2 CC:3 DD:4 EE:5]]
	fmt.Println("b6:", b6) // 打印结果: [map[AA:300 BB:2 CC:3 DD:4 EE:5] map[AA:1 BB:2 CC:3 DD:4 EE:5] map[AA:1 BB:2 CC:3 DD:4 EE:5]]
	fmt.Println("")

}

/*
a: [map[AA:300 BB:2 CC:3 DD:4 EE:5] map[AA:1 BB:2 CC:3 DD:4 EE:5]]
b: [map[AA:300 BB:2 CC:3 DD:4 EE:5] map[AA:1 BB:2 CC:3 DD:4 EE:5]]

a1: [map[AA:100 BB:500 CC:300 DD:4 EE:5] map[AA:1 BB:2 CC:3 DD:4 EE:5]]
b1: map[AA:100 BB:500 CC:300 DD:4 EE:5]

a2: [map[AA:300 BB:2 CC:3 DD:4 EE:5] map[AA:1 BB:2 CC:3 DD:4 EE:5]]
b2: map[AA:300 BB:2 CC:3 DD:4 EE:5]

a3: [map[AA:300 BB:2 CC:3 DD:4 EE:5] map[AA:1 BB:2 CC:3 DD:4 EE:5]]
b3: map[AA:300 BB:2 CC:3 DD:4 EE:5]

a4: [map[AA:300 BB:2 CC:3 DD:4 EE:5] map[AA:1 BB:2 CC:3 DD:4 EE:5]]
b4: [map[AA:300 BB:2 CC:3 DD:4 EE:5] map[AA:1 BB:2 CC:3 DD:4 EE:5]]

a5: [map[AA:300 BB:2 CC:3 DD:4 EE:5] map[AA:1 BB:2 CC:3 DD:4 EE:5]]
b5: [map[AA:300 BB:2 CC:3 DD:4 EE:5] map[AA:1 BB:2 CC:3 DD:4 EE:5] map[AA:1 BB:2 CC:3 DD:4 EE:5]]

a6: [map[AA:1 BB:2 CC:3 DD:4 EE:5] map[AA:1 BB:2 CC:3 DD:4 EE:5]]
b6: [map[AA:300 BB:2 CC:3 DD:4 EE:5] map[AA:1 BB:2 CC:3 DD:4 EE:5] map[AA:1 BB:2 CC:3 DD:4 EE:5]]
*/

// Modify9 函数，直接操作，会影响原参数
// Modify10函数，单独操作里面的map元素，会影响原参数
// Modify11函数，取址操作里面的map元素，会影响原参数
// Modify12函数，单独操作里面的map元素，会影响原参数
// Modify13函数，copy操作后操作里面的map元素，会影响原参数
// Modify14函数，append操作后操作里面的map元素，会影响原参数
// Modify15函数，append操作后,深度复制map元素，再操作深度复制后的map元素，不会影响原参数

// 一句话总结，slice操作是传址，append后传值 map的操作都是传址
