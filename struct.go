package main

import "fmt"

func change(temp *Student, key string , value string){
	if key == "name" {
		temp.name=value
	}else if key=="branch"{
		temp.branch=value
	}else if key=="college"{
		temp.college=value
	}
}

func (x *Student) change(value int){
	x.roll=value
}
type Student struct{
	name string
	roll int
	branch string
	college string
}

func main(){
	a:=Student{"Prajwal",92,"Computer science","Ramaiah"}
	fmt.Println(a)
	change(&a,"college","Ramaiah Institute of technology")
	a.change(44)
	fmt.Println(a)
	b:=&Student{"Rahul",77,"Civil","Nagarjuna"}
	b.college="Nagarjuna Institute of technology"
	fmt.Println(*b)
}
