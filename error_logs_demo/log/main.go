package main

import (
	"log"
)

func main() {
	//log.Print("Hey,I'm a log")//2021/04/13 00:43:59 Hey,I'm a log

	//log.Fatal("Hey,I'm an error log")//2021/04/13 00:45:18 Hey,I'm an error log,记录错误并结束程序
	//fmt.Print("Can you see me?")

	//log.Panic("Hey,I'm an error log!")// 停止程序，并打印错误堆栈信息
	//fmt.Print("Can you see me?")

	log.SetPrefix("main(): ") //想程序的日志消息添加前缀
	log.Print("Hey,I'm a log!")
	log.Fatal("Hey,I'm an error log!")

}
