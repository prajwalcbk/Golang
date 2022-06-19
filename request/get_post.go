package main

import ("fmt"
	"net/http"
	"io/ioutil"
	"os"
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

func ff(data string){
	fmt.Print(data)
}


func make_urlValues(param string) url.Values{
	s:=strings.Split(param,";")
	q:=url.Values{}
	for _,dd := range s {
		key_value:=strings.Split(dd,"=")
		q.Add(key_value[0],key_value[1])
	}
	return q
	}

func make_usn_urlValues(param string,usnn string) url.Values{
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

func getmsritResult(method string,urll string,cookie string,param string,usn string){
	client := http.Client{}
	data:=make_usn_urlValues(param,usn)
	req,err:=http.NewRequest(method,urll,strings.NewReader(data.Encode()))
	print_err(err)
	req.Header.Set("Cookie",cookie)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res , err := client.Do(req)
	print_err(err)
	data_b,err:=ioutil.ReadAll(res.Body)
	print_err(err)
	//ff(string(data_b))
	doc, err := htmlquery.Parse(strings.NewReader(string(data_b)))
	ff(get_data(doc, "/html/body/div[2]/div/div/div[4]/div/div/div[1]/div/div[4]/div/p"))
	ff(" "+get_data(doc,"/html/body/div[2]/div/div/div[3]/div/div/div[2]/div/h2"))
	ff(" "+get_data(doc,"/html/body/div[2]/div/div/div[3]/div/div/div[1]/div/h3"))
	fmt.Println()
}

func postRequest(method string,urll string,cookie string,param string){
	client := http.Client{}
	data:=make_urlValues(param)
	req,err:=http.NewRequest(method,urll,strings.NewReader(data.Encode()))
	print_err(err)
	req.Header.Set("Cookie",cookie)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res , err := client.Do(req)
	print_err(err)
	data_b,err:=ioutil.ReadAll(res.Body)
	print_err(err)
	fmt.Println(string(data_b))
}


func getRequest(method string,urll string,cookie string,param string){
	client := http.Client{}
	req,err:= http.NewRequest(method,urll,nil)
	print_err(err)
	req.Header.Set("Cookie",cookie)
	q:=make_urlValues(param)
	req.URL.RawQuery=q.Encode()
	res , err := client.Do(req)
	print_err(err)
	data,err:=ioutil.ReadAll(res.Body)
	print_err(err)
	fmt.Println(string(data))
}
func input(content string){
	fmt.Print("Enter "+content+" > ")
}
func main(){
		if(len(os.Args)==1){

			var m,u,c,d,us string
			var msrit bool
			var t int
			fmt.Println("MSRIT Results true/false")
			fmt.Scanln(&msrit)
			input("Request method GET OR POST ")
			fmt.Scanln(&m)
			input("Target URL")
			fmt.Scanln(&u)
			input("Cookie")
			fmt.Scanln(&c)
			input("Form Data or Get parameters")
			fmt.Scanln(&d)

			if(msrit) {
			input("Starting USN")
			fmt.Scanln(&us)
			input("Total Requests")
			fmt.Scanln(&t)

			for i:=1;i<t;i++ {
				usn:=""
				if i < 10 {
					usn=us+"00"+strconv.Itoa(i)
				}else if i<100 {
					usn=us+"0"+strconv.Itoa(i)
				}else{
					usn=us+strconv.Itoa(i)
				}
				getmsritResult(m,u,c,d,usn)
			}
		}else if (m=="GET"){
			getRequest(m,u,c,d)
		} else {
			postRequest(m,u,c,d)
		}
		os.Exit(0)
	}else if (len(os.Args)!=5 && len(os.Args)!=7) {
			fmt.Println("\n\nUsage of program is $go run get2cookie_commandline.go <method> <target_url> <cookie> <data>")
			fmt.Print("Example > ")
			fmt.Println("	$go run get2cookie_commandline.go GET http://172.17.0.2 \"PHPSESSID=h0o8ajpmmno1ul8b6rrjv3fpj1;security=low\" \"id=1;Submit=Submit\"\n")
				fmt.Println("Usage of program for MSRIT Results is $go run get2cookie_commandline.go <method> <target_url> <cookie> <data> <startusn> <total>")
			fmt.Print("Example > ")
			fmt.Println("	$go run get2cookie_commandline.go GET http://172.17.0.2 \"PHPSESSID=h0o8ajpmmno1ul8b6rrjv3fpj1;security=low\" \"id=1;Submit=Submit\" \"1MS18CS\" 142\n\n")

			os.Exit(2)
		}
		method:=os.Args[1]
		urll:=os.Args[2]
		cookie:=os.Args[3]
		param:=os.Args[4]
		if (method=="GET" && len(os.Args)==5){
			getRequest(method,urll,cookie,param)
		} else if(len(os.Args)==5){
			postRequest(method,urll,cookie,param)
		}else if(len(os.Args)==7){
			startusn:=os.Args[5]
			maxi,_:=strconv.Atoi(os.Args[6])
			for i:=1;i<maxi;i++ {
				usn:=""
				if i < 10 {
					usn=startusn+"00"+strconv.Itoa(i)
				}else if i<100 {
					usn=startusn+"0"+strconv.Itoa(i)
				}else{
					usn=startusn+strconv.Itoa(i)
				}
				getmsritResult(method,urll,cookie,param,usn)
			}
		}

	}
