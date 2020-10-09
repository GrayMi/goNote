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

func Modify9(a []ST) []ST {
	a[0].AA = 300
	return a
}

func Modify10(a []ST) ST {
	b := a[0]
	b.AA = 100
	b.BB = 500
	b.CC = 300
	return b
}

func Modify11(a []ST) ST {
	b := new(ST)
	b = &a[0]
	b.AA = 300
	return *b
}

func Modify12(a []ST) ST {
	b := ST{}
	b = a[0]
	b.AA = 300
	return b
}

func Modify13(a []ST) []ST {
	b := make([]ST, 2)
	copy(b, a)
	b[0].AA = 300
	return b
}

func Modify14(a []ST) []ST {
	d := ST{
		AA: 1,
		BB: 2,
		CC: 3,
		DD: 4,
		EE: 5,
	}
	a[0].AA = 100
	b := append(a, d)
	b[1].AA = 300
	return b
}

func main() {
	a := []ST{
		{
			AA: 1,
			BB: 2,
			CC: 3,
			DD: 4,
			EE: 5,
		},
		{
			AA: 1,
			BB: 2,
			CC: 3,
			DD: 4,
			EE: 5,
		},
	}
	b := Modify9(a)
	fmt.Println("a:", a) // 打印结果:[{300 2 3 4 5} {1 2 3 4 5}]
	fmt.Println("b:", b) // 打印结果:[{300 2 3 4 5} {1 2 3 4 5}]
	fmt.Println("")

	a1 := []ST{
		{
			AA: 1,
			BB: 2,
			CC: 3,
			DD: 4,
			EE: 5,
		},
		{
			AA: 1,
			BB: 2,
			CC: 3,
			DD: 4,
			EE: 5,
		},
	}
	b1 := Modify10(a1)
	fmt.Println("a1:", a1) // 打印结果:[{1 2 3 4 5} {1 2 3 4 5}]
	fmt.Println("b1:", b1) // 打印结果:{100 500 300 4 5}
	fmt.Println("")

	a2 := []ST{
		{
			AA: 1,
			BB: 2,
			CC: 3,
			DD: 4,
			EE: 5,
		},
		{
			AA: 1,
			BB: 2,
			CC: 3,
			DD: 4,
			EE: 5,
		},
	}
	b2 := Modify11(a2)
	fmt.Println("a2:", a2) // 打印结果:[{300 2 3 4 5} {1 2 3 4 5}]
	fmt.Println("b2:", b2) // 打印结果:{300 2 3 4 5}
	fmt.Println("")

	a3 := []ST{
		{
			AA: 1,
			BB: 2,
			CC: 3,
			DD: 4,
			EE: 5,
		},
		{
			AA: 1,
			BB: 2,
			CC: 3,
			DD: 4,
			EE: 5,
		},
	}
	b3 := Modify12(a3)
	fmt.Println("a3:", a3) // 打印结果:[{1 2 3 4 5} {1 2 3 4 5}]
	fmt.Println("b3:", b3) // 打印结果:{300 2 3 4 5}
	fmt.Println("")

	a4 := []ST{
		{
			AA: 1,
			BB: 2,
			CC: 3,
			DD: 4,
			EE: 5,
		},
		{
			AA: 1,
			BB: 2,
			CC: 3,
			DD: 4,
			EE: 5,
		},
	}
	b4 := Modify13(a4)
	fmt.Println("a4:", a4) // 打印结果:[{1 2 3 4 5} {1 2 3 4 5}]
	fmt.Println("a4:", b4) // 打印结果:[{300 2 3 4 5} {1 2 3 4 5}]
	fmt.Println("")

	a5 := []ST{
		{
			AA: 1,
			BB: 2,
			CC: 3,
			DD: 4,
			EE: 5,
		},
		{
			AA: 1,
			BB: 2,
			CC: 3,
			DD: 4,
			EE: 5,
		},
	}
	b5 := Modify14(a5)
	fmt.Println("a5:", a5) // 打印结果:[{100 2 3 4 5} {1 2 3 4 5}]
	fmt.Println("b5:", b5) // 打印结果:[{100 2 3 4 5} {300 2 3 4 5} {1 2 3 4 5}]
	fmt.Println("")

}

/*
a: [{300 2 3 4 5} {1 2 3 4 5}]
b: [{300 2 3 4 5} {1 2 3 4 5}]

a1: [{1 2 3 4 5} {1 2 3 4 5}]
b1: {100 500 300 4 5}

a2: [{300 2 3 4 5} {1 2 3 4 5}]
b2: {300 2 3 4 5}

a3: [{1 2 3 4 5} {1 2 3 4 5}]
b3: {300 2 3 4 5}

a4: [{1 2 3 4 5} {1 2 3 4 5}]
a4: [{300 2 3 4 5} {1 2 3 4 5}]

a5: [{100 2 3 4 5} {1 2 3 4 5}]
b5: [{100 2 3 4 5} {300 2 3 4 5} {1 2 3 4 5}]
*/

// Modify9 函数，直接操作，会影响原参数
// Modify10函数，单独操作里面的struct元素，不会影响原参数
// Modify11函数，取址操作，会影响原参数
// Modify12函数，单独操作里面的struct元素，不影响原参数
// Modify13函数，copy后操作，不影响原参数
// Modify13函数，append后操作，不影响原参数
