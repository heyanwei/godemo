package main

import "log"

func main() {
	//创建a的指针
	var a *int
	b := 12
	//把b的地址赋值给a
	a = &b
	//打印a的值
	log.Println(*a)
	//打印a的地址
	log.Println(a)
}
