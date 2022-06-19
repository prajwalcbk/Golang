package main
import "fmt"


type Student struct{
	name string
	usn int
}

func main(){
	var arr [5]int
	for i:=0 ; i<5;i++ {
		arr[i]=i
	}
	for i:=0 ; i<5 ;i++ {
		//fmt.Print(arr[i]," ")
		fmt.Println(arr[i])
	}
	var count int
	fmt.Println("Enter the Number of Students")
	fmt.Scanln(&count)
	var list [3]Student
	for i:=0;i<count;i++ {
		fmt.Println("Enter the Details of ",i+1," Student")
		fmt.Print("Name >> ")
		fmt.Scanln(&list[i].name)
		fmt.Print("USN >> ")
		fmt.Scanln(&list[i].usn)
	}
	for i:=0 ; i< count ; i++ {
		fmt.Println(list[i])
	}

}
