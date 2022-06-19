package main

import ("fmt"
	"net/http"
	"io/ioutil"
)

func print_err(err error){
	if err!=nil {
		fmt.Println(err)
	}
}
func getResult(urll string){
	client := http.Client{}
	req ,err := http.NewRequest("GET", urll, nil)
	print_err(err)
	req.Header.Set("Cookie","PHPSESSID=h0o8ajpmmno1ul8b6rrjv3fpj1; security=low")
	res , err := client.Do(req)
	fmt.Println(res)
	print_err(err)
	data,err:=ioutil.ReadAll(res.Body)
	print_err(err)
	fmt.Println(string(data))
}
func main(){
		urll:="http://172.17.0.2/index.php"
		getResult(urll)
	}
