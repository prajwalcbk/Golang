package main

import ( "fmt"
	"os"
	"net/http"
	"io"
)

func print_err(err error){
	if err!=nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func main(){
	//fmt.Println("")
	url:=os.Args[1]
	res,err:=http.Get(url)
	print_err(err)
	_,err=io.Copy(os.Stdout,res.Body)
	print_err(err)
}

