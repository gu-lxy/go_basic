package main

import "fmt"

func main() {
		//变量的声明
		//var  变量名  变量类型
	var	a int
	fmt.Println(a)
	//批量声明
	var (
		b int
		c string
		d []float64
		e func() bool
		f struct{
			x int
			y string
		}
	)
	fmt.Println(b,c,d,e,f)
	//变量初始化
	//var 变量名 变量类型 = 表达式
	var a1 int = 10
	fmt.Println(a1)
	// var 变量名 = 表达式，编译器自动推断类型
	var a2 = 10
	fmt.Println(a2)
	//变量名 := 表达式
	a3 := 10
	fmt.Println(a3)
	//重复定义变量a4
	var a4 = 20
	fmt.Println(a4)
	a4 = 10
	fmt.Println(a4)//编译器会报错
	//可如下操作
	var a5 = 10
	fmt.Println(a5)
	a5,b1 := 20,30
	fmt.Println(a5,b1)
	//变量多重复赋值
	var ax int = 20
	var bx int = 10
	var ser int
	ser = ax
	ax = bx
	fmt.Println(ax,bx,ser)
	//
	var ay int = 10
	var by int = 20
	fmt.Println(a,b)
	ay = a ^ b
	fmt.Println(a)
	by = b ^ a
	fmt.Println(b)
	ay = a ^ b
	fmt.Println(ay,by)

	a, _ = GetData() //舍弃第二个返回值
	_, b = GetData() //舍弃第一个返回值
	fmt.Println(a,b)
}
//匿名变量
func GetData() (int,int) {
	return 10, 20
}
