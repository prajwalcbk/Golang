package main
import ( "fmt"
	"time"
	"strconv"
	"os"
)
func main(){
	start := time.Now()
	s:=os.Args[1]
	rang ,err := strconv.Atoi(s)
	if err != nil{
		fmt.Println(err)
		os.Exit(2)
	}
	for i:=0 ; i<rang ; i+=1 {
		continue
	}
	elapsed := time.Since(start)
	fmt.Println("Total time taken to execute "+s+" Iterations is ", elapsed)
}
