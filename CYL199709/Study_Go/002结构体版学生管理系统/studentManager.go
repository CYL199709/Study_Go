package main
import "fmt"

type student struct{
	id int64
	name string
}

//造一个学生的管理者
type studentManager struct{
	allStudent map[int64]student
}

//查看学生
func (s studentManager)showStudents(){
	for _,stu := range(s.allStudent){
		fmt.Printf("学号：%d 姓名：%s \n",stu.id,stu.name)
	}

}

//增加学生
func (s studentManager)addStudent(){
	var(
		stuID int64
		stuName string
	)
	fmt.Print("请输入学号：")
	fmt.Scanln(&stuID)
	fmt.Print("请输入姓名：")
	fmt.Scanln(&stuName)
	newStu := student{
		id: stuID,
		name: stuName,
	}
	s.allStudent[newStu.id] = newStu
	fmt.Println("添加成功！")
	
}

//修改学生
func (s studentManager)editStudent(){
	var stuID int64
	fmt.Print("请输入学号：")
	fmt.Scanln(&stuID)
	stuObj, ok := s.allStudent[stuID]
	if !ok{
		fmt.Println("查无此人！")
		return
	}
	fmt.Printf("你要修改的学生信息如下：学号 %d 姓名 %s\n",stuObj.id,stuObj.name)
	fmt.Print("请输入新学生的名字：")
	var newName string
	fmt.Scanln(&newName)
	stuObj.name = newName
	s.allStudent[stuID] = stuObj
}

//删除学生
func (s studentManager)deleteStudent(){
	var stuID int64
	fmt.Print("请输入要删除的的学生学号：")
	fmt.Scanln(&stuID)
	stuObj, ok := s.allStudent[stuID]
	if !ok{
		fmt.Println("查无此人！")
		return
	}
	fmt.Printf("你要删除的学生信息如下：学号 %d 姓名 %s\n",stuObj.id,stuObj.name)
	delete(s.allStudent,stuID)
	fmt.Println("删除成功！")
}