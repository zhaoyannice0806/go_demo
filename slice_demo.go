package main

import "fmt"

func testSliceDemo() {
	fmt.Println("===============SLICE 声明=================")
	var intSlice []int
	var strSlice []string
	fmt.Printf("str slice len:%d,cap=%d,值：%v\n", len(strSlice), cap(strSlice), strSlice)
	fmt.Printf("int slice len:%d,cap=%d,值：%v\n", len(intSlice), cap(intSlice), intSlice)

	fmt.Printf("\n=======SLICE make初建 指定长度和容量========\n")
	slice1 := make([]int, 5, 10)
	slice2 := make([]string, 2)
	fmt.Printf("islice1 len:%d,cap=%d,值：%v\n", len(slice1), cap(slice1), slice1)
	fmt.Printf("slice2 len:%d,cap=%d,值：%v\n", len(slice2), cap(slice2), slice2)

	fmt.Printf("\n=====SLICE 从已有数组中或者切片中截取片段======\n")
	int1 := [5]int{1, 2, 3, 4, 5}

	// 截取数组 1到3，不包含3
	s1 := int1[1:3]
	fmt.Printf("s1 len:%d,cap=%d,值：%v\n", len(s1), cap(s1), s1)

	// 截取数组 0到3，包含3
	s2 := int1[:3]
	fmt.Printf("s2 len:%d,cap=%d,值：%v\n", len(s2), cap(s2), s2)

	// 截取数组 1到5，包含5
	s3 := int1[1:]
	fmt.Printf("s3 len:%d,cap=%d,值：%v\n", len(s3), cap(s3), s3)

	// 截取数组 0到5，包含5
	s4 := int1[:]
	fmt.Printf("s4 len:%d,cap=%d,值：%v\n", len(s4), cap(s4), s4)

	// 从slice截取
	s5 := s1[1:2]
	fmt.Printf("s5 len:%d,cap=%d,值：%v\n", len(s5), cap(s5), s5)

	fmt.Printf("\n===================SLICE 核心操作===============\n")
	fmt.Printf("\n===================SLICE 添加元素===============\n")

	// 添加元素
	s6 := make([]int, 2, 4)
	s6[0] = 0
	s6[1] = 1
	fmt.Printf("s5 len:%d,cap=%d,值：%v\n", len(s6), cap(s6), s6)

	// 追加一个元素
	s6 = append(s6, 2)
	fmt.Printf("s5 len:%d,cap=%d,值：%v\n", len(s6), cap(s6), s6)

	// 追加多个元素
	s6 = append(s6, 2, 3)
	fmt.Printf("s5 len:%d,cap=%d,值：%v\n", len(s6), cap(s6), s6)

	// 追加一个切片
	s7 := []int{4, 5, 6, 7, 8}
	s6 = append(s6, s7...)
	fmt.Printf("s5 len:%d,cap=%d,值：%v\n", len(s6), cap(s6), s6)

	fmt.Printf("\n===================SLICE 遍历===============\n")
	// for 循环
	for i := 0; i < len(s6); i++ {
		fmt.Println(s6[i])
	}

	// range 循环
	for i, v := range s6 {
		fmt.Println(i, v)
	}

	// for 循环 忽略索引
	for _, v := range s6 {
		fmt.Println(v)
	}

	fmt.Printf("\n===================SLICE 修改元素===============\n")
	// 切片是引用类型，直接通过索引修改元素会同步到底层数组（共享底层数组的切片也会受影响）。
	arr := [5]int{1, 2, 3, 4, 5}
	s8 := arr[:]
	s9 := s8[1:]

	s8[2] = 100
	fmt.Printf("arr=%v \n", arr)
	fmt.Printf("s8=%v \n", s8)
	fmt.Printf("s9=%v \n", s9)

	fmt.Printf("\n===================SLICE 拷贝===============\n")
	s10 := make([]int, 2)
	copy(s10, s6)
	fmt.Printf("s10=%v \n", s10)

	s11 := make([]int, 20)
	copy(s11, s6)
	fmt.Printf("s11=%v \n", s11)

	fmt.Printf("\n===================SLICE 删除元素===============\n")
	s12 := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", s12[:2])
	fmt.Printf("%v\n", s12[3:])
	s12 = append(s12[:2], s12[3:]...)
	fmt.Printf("s12=%v \n", s12)

}
