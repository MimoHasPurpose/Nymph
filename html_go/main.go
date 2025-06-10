package main
import (
  
    "os"
	"math/rand/v2"
	"strconv"
	"fmt"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}
// function adds html, css and javascript files.
func createfile(){
	a:=rand.IntN(100)
    ht := []byte("<head>\n<title></title>\n</head>\n<body>\n</body>")
	cs := []byte("*{background-color: black; size:20px;}")
	js := []byte("console.log(\"hello world)")
    t:=strconv.Itoa(a)
	fmt.Printf("what file do u wanna create!\n 1. html \n2.css \n3.javascript\n  ")
	var i int
	fmt.Scan(&i)
	switch i {
    case 1:
		err := os.WriteFile("E:\\Github\\Nymph\\html_go\\tmp\\"+"file"+t+".html", ht, 0644)
    	check(err)
    case 2:
        err := os.WriteFile("E:\\Github\\Nymph\\html_go\\tmp\\"+"file"+t+".css", cs, 0644)
    	check(err)
    case 3:
        err := os.WriteFile("E:\\Github\\Nymph\\html_go\\tmp\\"+"file"+t+".js", js, 0644)
    	check(err)
	case 4:
		err := os.WriteFile("E:\\Github\\Nymph\\html_go\\tmp\\"+"file"+t+".html", ht, 0644)
		err := os.WriteFile("E:\\Github\\Nymph\\html_go\\tmp\\"+"file"+t+".css", cs, 0644)
		err := os.WriteFile("E:\\Github\\Nymph\\html_go\\tmp\\"+"file"+t+".js", js, 0644)
		check(err)
    }
}

func main() {
	
	createfile();
	
    

}