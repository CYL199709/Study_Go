package main
import "fmt"
//	map
func main(){
	var m1 map[string]int
	m1 = make(map[string]int,10) //初始化，要估算好该map容量，避免在程序运行期间再动态扩容
	m1["微信"] = 1
	m1["支付宝"] = 2

	fmt.Println(m1)

	value, ok := m1["淘宝"]
	if !ok{
		fmt.Println("查无此key")
	} else{
		fmt.Println(value)
	}
	//删除
	delete(m1,"微信")
	//看内置函数文档 go doc builtin.delete
	//网址：studygolang.com/pkgdoc
	fmt.Println(m1)

	
}