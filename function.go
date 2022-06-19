package main 

import "fmt"

func add(a,b int) int  {
	return a+b
}

func welcome(name string){
	fmt.Println("Welcome ",name)
}
func main(){
	welcome("Prajwal")
	welcome("Prajjuu!!!")
	welcome("Amuuu")
	fmt.Println(add(5,6))
}
