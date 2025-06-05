package main

import (
  
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}
func project(project_name){
	err:=os.Mkdir("${project_name}",0755)
	check(err)
}
func file(file_name){
	create:=func(file_name string){
		d:=[]byte("")
		check(os.WriteFile(file_name,d,0644))
	}
	create("{project_name}/{file_name}")
}
func main() {
	
    d1 := []byte("<head>\n<title></title>\n</head>\n<body>\n</body>")
    err := os.WriteFile("E:\\Github\\Nymph\\html_go\\tmp\\basic.html", d1, 0644)
    check(err)
    

}