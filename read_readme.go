package main

import "fmt"
import "os"

func main(){
	data, error := os.ReadFile("README.md")
	if(error != nil){
		panic(error)
	}

	// convert to string because ASCII default
	fmt.Println(string(data))

}