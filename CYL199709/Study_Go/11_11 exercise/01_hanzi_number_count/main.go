package main
import "unicode"
import "fmt"
func main(){
	//1.判断字符串中汉字的数量
	s1 := "Hello沙河"
	var count int
	for _,c := range s1{
		if unicode.Is(unicode.Han,c){
			count++
		}
	} 
	fmt.Println(count)


}