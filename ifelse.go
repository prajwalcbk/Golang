package main 
import "fmt"

func main(){
	var age int
	fmt.Println("Enter the age ")
	fmt.Scanln(&age)
	if age>18 {
		fmt.Println("Eligible to vote")
	}else{
		fmt.Println("Not Eligible to vote")
	}
}
