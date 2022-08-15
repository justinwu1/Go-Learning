package main

import (
	"fmt"
	"net/http"
)

// Create different routes
func Home(writer http.ResponseWriter,request *http.Request){
	n,err := fmt.Fprintf(writer,"Home Page")
	if err != nil{
		fmt.Println("there's an error")
	}
	fmt.Println(fmt.Sprintf("The # of bytes %d",n ))
}

func About(writer http.ResponseWriter, request * http.Request){
	sum := addValues(1,1)
	_,err := fmt.Fprintf(writer,fmt.Sprintf("This is about page and addValues() returned %d",sum ))
	if err != nil{
		fmt.Println("there's an error")
	}
}

func addValues(x int, y int) int{
	return x + y
}

func main() {
	http.HandleFunc("/",Home)
	http.HandleFunc("/About",About)
	_ = http.ListenAndServe(":8080",nil)

}