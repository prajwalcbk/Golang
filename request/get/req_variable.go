package main

import ( "fmt"
	"os"
	"net/http"
	"io/ioutil"
)
func print_err(err error){
	if err!=nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func main(){
	url:=os.Args[1]
	res,err:=http.Get(url)
	print_err(err)
	body,err:=ioutil.ReadAll(res.Body)
	print_err(err)
	bbody:=string(body)
	fmt.Println(bbody)
}

