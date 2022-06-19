package main
import ("fmt"
	"encoding/json"
	"net/url"
	"net/http"
	"log"
)

func getResult(urll string , data , cookie url.Values){
	client := http.Client{}
	req , err := http.NewRequest("POST", urll, data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header = http.Header{
		"Content-Type": {"application/json"},
		"Cookie":cookie,
	}
	res , err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
    var res map[string]interface{}
    fmt.Println(res.Body)
    json.NewDecoder(res.Body).Decode(&res)

    fmt.Println(res["form"])
}
func main(){
	for i:=0; i<142;i+=1 {
		urll:="http://exam.msrit.edu/index.php"
		var ii string
		if i <10 {
			ii="00"+string(i)
		}else if i<100 {
			ii="0"+string(i)
		}else {
			ii=string(i)
		}
		usn:="1MS18CS"+ii
		cookie := url.Values{
"3e84ce0de6b7d1eb79699e2a6adfb3a1":	{"50c6ckm9suo4ocg12hdfgvr6l6"},
			"_ga":		{"GA1.2.261898089.1649329895"},
			"_gid":		{"GA1.2.1337257613.1650266766"},
		}
		data := url.Values{
			"usn":{usn},
			"osolCatchaTxt":{"NG94M"},
			"osolCatchaTxtInst":{"0"},
			"option":{"com_examresult"},
			"task":{"getResult"},
		}
		getResult(urll,data,cookie)
	}
}
