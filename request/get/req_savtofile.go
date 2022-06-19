package main

import ( "fmt"
	"os"
	"net/http"
)

func print_err(err error){
	if err!=nil {
		fmt.Println(err)
	}
}



func main(){
	url:=os.Args[1]
	res,err:=http.Get(url)
	print_err(err)
	f,err:=os.Create("index.html")
	print_err(err)
	_,err=f.ReadFrom(res.Body)
	print_err(err)
}

