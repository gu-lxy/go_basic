package main

import "fmt"

func main() {
	//整形
	var a int8//声明有符号8位整形[-2^8~2^8)
	var a uint8//声明无符号8位整形[0~2^8)
	//浮点型
	var x float32//声明32位浮点数
	//复数型
	complex64()
	complex128()
	//字符串
	var atr string
	atr = "hello"
	fmt.Println(atr)
	var temp string
	temp = `
		x := 10
		y := 20
		z := 30
		fmt.Println(x,"   ",y,"  ", z)
		x, y, z = y, z, x
	    fmt.Println(x,"   ",y,"  ", z)`
	fmt.Println(temp)
	//布尔型
	var flag bool
}
