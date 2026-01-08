package main

import (
	"fmt"
	"strconv"
)

func testTypeConversionDemo() {
	var a int32 = 13
	var b byte = 3

	// 数字类型可以直接强制转换
	var f float32 = float32(a) / float32(b)
	fmt.Println(f)

	//当int32类型转为byte类型 高位丢失
	var i1 int32 = 256
	var b1 = byte(i1)
	fmt.Println(b1)

	/**
	byte 类型
		uint8 的别名 - 表示8位无符号整数（0-255）
		单字节存储 - 每个值恰好存储一个字节的数据
		适用于ASCII字符 - 适合处理ASCII字符和二进制数据
		固定大小 - 始终是1字节/值
	rune 类型
		int32 的别名 - 表示32位有符号整数
		Unicode码点 - 可以存储任何Unicode字符
		可变字节大小 - 在UTF-8中可以表示1-4字节的字符
		支持国际字符 - 能处理中文、日文、阿拉伯文等

	byte 类型
		uint8 的别名 - 表示8位无符号整数（0-255）
		单字节存储 - 每个值恰好存储一个字节的数据
		适用于ASCII字符 - 适合处理ASCII字符和二进制数据
		固定大小 - 始终是1字节/值
	rune 类型
		int32 的别名 - 表示32位有符号整数
		Unicode码点 - 可以存储任何Unicode字符
		可变字节大小 - 在UTF-8中可以表示1-4字节的字符
		支持国际字符 - 能处理中文、日文、阿拉伯文等

	*/
	str := "hello world 123 你好"
	var bytes []byte = []byte(str)
	var runes []rune = []rune(str)
	fmt.Printf("bytes : %v \n", bytes)
	fmt.Printf("runes : %v \n", runes)

	str2 := string(bytes)
	str3 := string(runes)
	fmt.Printf("str2 : %s \n", str2)
	fmt.Printf("str3 : %s \n", str3)

	str4 := "123"

	// 字符串转数字
	i, err := strconv.Atoi(str4)
	fmt.Printf("数字转字符串结果：%d \n", i)
	if err != nil {
		panic(err)
	}

	// 数字转字符串
	str5 := strconv.Itoa(i)
	fmt.Printf("字符串转数字结果：%s \n", str5)

	// 转10进制
	ui1, err := strconv.ParseUint(str5, 10, 32)
	if err == nil {
		fmt.Printf("字符串转十进制数字结果：%d \n", ui1)
	}

	ui2, err := strconv.ParseUint("FF", 16, 32)
	if err == nil {
		fmt.Printf("十六进制字符串转数字结果：%d \n", ui2)
	}

	ui3, err := strconv.ParseUint("1010", 2, 32)
	if err == nil {
		fmt.Printf("二进制字符串转数字结果：%d \n", ui3)
	}

	str6 := strconv.FormatInt(int64(i), 2)
	fmt.Printf("数字转二进制字符串结果：%s \n", str6)

	str7 := strconv.FormatInt(int64(i), 8)
	fmt.Printf("数字转八进制字符串结果：%s \n", str7)

	str8 := strconv.FormatInt(int64(i), 16)
	fmt.Printf("数字转十六进制字符串结果：%s \n", str8)
}
