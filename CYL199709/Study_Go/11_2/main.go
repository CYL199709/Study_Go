package main

import "fmt"

func main() {
	s2 := "白萝卜"      //按asicc位存储，无法直接修改，每个中文字符3-4个字节
	s3 := []rune(s2) //把字符串强制转换成切片
	s3[0] = '红'
	fmt.Println(string(s3))

}
