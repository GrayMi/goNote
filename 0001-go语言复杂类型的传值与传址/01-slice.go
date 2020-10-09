package main

import (
	"fmt"
)

// 先改变slice元素，再增加长度 ---- 原slice元素改变,新slice改变
func Modify3(array []int) []int {
	array[0] = 300
	t1 := append(array, 6)
	return t1
}

// 先改变slice元素，再缩小长度 ---- 原slice元素改变,新slice改变
func Modify4(array []int) []int {
	array[0] = 300
	t1 := array[0:3]
	return t1
}

// 先增加长度，再改变新slice元素值 ---- 原slice元素不改变，新slice改变
func Modify5(array []int) []int {
	t1 := append(array, 6)
	t1[0] = 300
	return t1
}

// 先增加长度，再改变原slice元素值 ---- 原slice元素改变，新slice不改变
func Modify6(array []int) []int {
	t1 := append(array, 6)
	array[0] = 300
	return t1
}

// 先缩小长度，再改变原slice元素值 ---- 原slice元素改变，新slice改变
func Modify7(array []int) []int {
	t1 := array[0:3]
	array[0] = 300
	return t1
}

// 先缩小长度，再改变新slice元素值 ---- 原slice元素改变，新slice改变
func Modify8(array []int) []int {
	t1 := array[0:3]
	t1[0] = 300
	return t1
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := Modify3(arr1)
	fmt.Println("arr1:", arr1) // 打印结果:[300 2 3 4 5]
	fmt.Println("arr2:", arr2) // 打印结果:[300 2 3 4 5 6]

	arr3 := []int{1, 2, 3, 4, 5}
	arr4 := Modify4(arr3)
	fmt.Println("arr3:", arr3) // 打印结果: [300 2 3 4 5]
	fmt.Println("arr4:", arr4) // 打印结果: [300 2 3]

	arr5 := []int{1, 2, 3, 4, 5}
	arr6 := Modify5(arr5)
	fmt.Println("arr5:", arr5) // 打印结果: [1 2 3 4 5]
	fmt.Println("arr6:", arr6) // 打印结果: [300 2 3 4 5 6]

	arr7 := []int{1, 2, 3, 4, 5}
	arr8 := Modify6(arr7)
	fmt.Println("arr7:", arr7) // 打印结果: [300 2 3 4 5]
	fmt.Println("arr8:", arr8) // 打印结果: [1 2 3 4 5 6]

	arr9 := []int{1, 2, 3, 4, 5}
	arr10 := Modify7(arr9)
	fmt.Println("arr9:", arr9)   // 打印结果: [300 2 3 4 5]
	fmt.Println("arr10:", arr10) // 打印结果: [300 2 3]

	arr11 := []int{1, 2, 3, 4, 5}
	arr12 := Modify8(arr11)
	fmt.Println("arr11:", arr11) // 打印结果: [300 2 3 4 5]
	fmt.Println("arr12:", arr12) // 打印结果: [300 2 3]
}

/*

$ go run main.go
arr1: [300 2 3 4 5]
arr2: [300 2 3 4 5 6]
arr3: [300 2 3 4 5]
arr4: [300 2 3]
arr5: [1 2 3 4 5]
arr6: [300 2 3 4 5 6]
arr7: [300 2 3 4 5]
arr8: [1 2 3 4 5 6]
arr9: [300 2 3 4 5]
arr10: [300 2 3]
arr11: [300 2 3 4 5]
arr12: [300 2 3]
*/

/*
```bash
|----------------------------------------------------------------------------|
|                                                                           |
| Modify3  先改变slice元素，再增加长度    ---- 原slice元素改变,新slice改变        |
| Modify4  先改变slice元素，再缩小长度    ---- 原slice元素改变,新slice改变        |
|                                                                           |
| Modify5  先增加长度，再改变新slice元素值 ---- 原slice元素不改变，新slice改变     |
| Modify6  先增加长度，再改变原slice元素值 ---- 原slice元素改变，新slice不改变     |
|                                                                           |
| Modify7  先缩小长度，再改变原slice元素值 ---- 原slice元素改变，新slice改变       |
| Modify8  先缩小长度，再改变新slice元素值 ---- 原slice元素改变，新slice改变       |
|                                                                           |
|---------------------------------------------------------------------------|

原slice append 扩容后，原始slice与新slice脱离关系
原slice 缩小后，原始slice与新slice保持关系
```
*/
