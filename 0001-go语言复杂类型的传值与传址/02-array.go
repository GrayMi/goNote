package main

import (
	"fmt"
)

// 数组array-值传递
func Modify1(array [5]int) [5]int {
	array[0] = 300
	return array
}

// slice-地址传递
func Modify2(array []int) []int {
	array[0] = 300
	return array
}

func main() {
	// 数组-需要明确指定大小
	arr5 := [5]int{1, 2, 3, 4, 5}
	arr6 := Modify1(arr5)
	fmt.Println("arr5:", arr5) // 打印结果:[1 2 3 4 5]
	fmt.Println("arr6:", arr6) // 打印结果:[300 2 3 4 5]

	// slice-不需要明确指定大小
	arr7 := []int{1, 2, 3, 4, 5}
	arr8 := Modify2(arr7)
	fmt.Println("arr7:", arr7) // 打印结果:[300 2 3 4 5]
	fmt.Println("arr8:", arr8) // 打印结果:[300 2 3 4 5]

	/*
		// 数组-大小不可以变化
		// 此处报错: ./main.go:32:16: first argument to append must be slice; have [5]int
		arr9 := append(arr5, 6)
		fmt.Println("arr9:", arr9) // 此处报错
	*/

	// slice-大小可以变化
	arr10 := append(arr7, 6)
	fmt.Println("arr10:", arr10) // 打印结果:[300 2 3 4 5 6]
}
