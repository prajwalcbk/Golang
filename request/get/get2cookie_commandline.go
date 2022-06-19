package main

import ("fmt"
	"net/http"
	"io/ioutil"
	"os"
//	"encoding/json"
//	"bytes"
	"strings"
	"net/url"
	"strconv"
	"golang.org/x/net/html"
	"github.com/antchfx/htmlquery"
)

func print_err(err error){
	if err!=nil {
		fmt.Println(err)
	}
}
func string_made(param string) url.Values{
		s:=strings.Split(param,";")
		q:=url.Values{}
		for _,dd := range s {
			key_value:=strings.Split(dd,"=")
			q.Add(key_value[0],key_value[1])
		}
		return q
	}

func string_data(param string,usnn string) url.Values{
		s:=strings.Split(param,";")
		q:=url.Values{}
		for _,dd := range s {
			key_value:=strings.Split(dd,"=")
			q.Add(key_value[0],key_value[1])
		}
		q.Add("usn",usnn)
		return q
	}

func get_data(doc *html.Node ,xp string) string {
list := htmlquery.Find(doc,xp)
	for _ , n := range list{
		return n.FirstChild.Data
	}
	return ""
}

func postResult(method string,urll string,cookie string,param string,startusn string){
	client := http.Client{}

	for i:=1;i<=142;i++ {
		data:=url.Values{}
		usn:=""
		if i < 10 {
			usn=startusn+"00"+strconv.Itoa(i)
		}else if i<100 {
			usn=startusn+"0"+strconv.Itoa(i)
		}else{
			usn=startusn+strconv.Itoa(i)
		}
	//fmt.Println(usn)
	data=string_data(param,usn)
	//fmt.Println(data)
	req,err:=http.NewRequest(method,urll,strings.NewReader(data.Encode()))
	print_err(err)
	req.Header.Set("Cookie",cookie)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res , err := client.Do(req)
	print_err(err)
	data_b,err:=ioutil.ReadAll(res.Body)
	print_err(err)
	//fmt.Println(string(data_b))
	doc, err := htmlquery.Parse(strings.NewReader(string(data_b)))
	fmt.Print(get_data(doc, "/html/body/div[2]/div/div/div[4]/div/div/div[1]/div/div[4]/div/p"))
	fmt.Print(" "+get_data(doc,"/html/body/div[2]/div/div/div[3]/div/div/div[2]/div/h2"))
	fmt.Print(" "+get_data(doc,"/html/body/div[2]/div/div/div[3]/div/div/div[1]/div/h3"))
	fmt.Println()
	}

}

func getResult(method string,urll string,cookie string,param string){
	client := http.Client{}
	req,err:= http.NewRequest(method,urll,nil)
	print_err(err)
	req.Header.Set("Cookie",cookie)
	q:=string_made(param)
	req.URL.RawQuery=q.Encode()
	res , err := client.Do(req)
	print_err(err)
	data,err:=ioutil.ReadAll(res.Body)
	print_err(err)
	fmt.Println(string(data))
}
func main(){
		if (len(os.Args)!=6) {
			fmt.Println("Usage of program is $go run get2cookie_commandline.go <target_url> <cookie>")
			fmt.Println("Example ")
			fmt.Println("$go run GET get2cookie_commandline.go http://172.17.0.2 \"PHPSESSID=h0o8ajpmmno1ul8b6rrjv3fpj1;security=low\" \"id=1;Submit=Submit\"")
			os.Exit(2)
		}
		method:=os.Args[1]
		urll:=os.Args[2]
		cookie:=os.Args[3]
		param:=os.Args[4]
		startusn:=os.Args[5]
		if method=="GET" {
			getResult(method,urll,cookie,param)
		} else {
			postResult(method,urll,cookie,param,startusn)
		}
	}
