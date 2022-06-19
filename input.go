package main

import "fmt"
func main(){
	var city string
	var i int
	fmt.Println("Enter the city name where u lived five years ago")
	fmt.Scanln(&i)
	fmt.Scanln(&city)
	fmt.Println("i =",i)
	fmt.Println("U lived in "+city+" 5 years ago")
}
