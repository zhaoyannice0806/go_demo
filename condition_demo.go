package main

import "fmt"

func testIfConditionDemo(a int) {
	if a > 5 {
		fmt.Println("大于5")
	} else if a >= 3 {
		fmt.Println("大于3")
	} else {
		fmt.Println("小于3")
	}

	var arr = []int{98, 56, 40, 87, 99, 100}
	for index, score := range arr {
		fmt.Printf("第%d位 成绩 %d", index+1, score)
		if score >= 90 {
			fmt.Printf("，结果 : 优秀\n")
		} else if score >= 80 {
			fmt.Printf("，结果 : 良好\n")
		} else if score >= 60 {
			fmt.Printf("，结果  : 及格\n")
		} else {
			fmt.Printf("，结果 : 不及格\n")
		}
	}
}

func testSwitchDemo(a int) {
	switch {
	case a > 5:
		fmt.Println("大于5")
	case a >= 3:
		fmt.Println("大于3")
	default:
		fmt.Println("小于3")
	}

	var arr = []int{98, 56, 40, 87, 99, 100}
	for index, score := range arr {
		fmt.Printf("第%d位 成绩 %d", index+1, score)
		switch {
		case score >= 90:
			fmt.Printf("，结果 : 优秀\n")
		case score >= 80:
			fmt.Printf("，结果 : 良好\n")
		case score >= 60:
			fmt.Printf("，结果 : 及格\n")
		default:
			fmt.Printf("，结果 : 不及格\n")
		}
	}

	// 基本switch
	days := []int{1, 3, 5, 7, 10}
	for index, day := range days {
		fmt.Printf("索引%d: %d\n", index, day)
		switch day {
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wednesday")
		case 4:
			fmt.Println("Thursday")
		case 5:
			fmt.Println("Friday")
		case 6, 7:
			fmt.Println("周末")
		default:
			fmt.Printf("Day %d: 未知\n", day)
		}
	}

	num := 1
	switch num {
	case 1:
		fmt.Println("执行 case 1")
		fallthrough // 穿透到 case 2
	case 2:
		fmt.Println("执行 case 2")
		fallthrough // 再次穿透到 case 3
	case 3:
		fmt.Println("执行 case 3")
		// 无 fallthrough，停止穿透
	default:
		fmt.Println("执行 default")
	}

}
