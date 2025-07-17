package main

import "fmt"

// net/http is the module that allows API calls
import "net/http"
import "io"
import "encoding/json"

func main(){
	// send a GET response to grab pikachu's data, and store the result in response and error
	response, error := http.Get("https://pokeapi.co/api/v2/pokemon/pikachu")

	if error != nil {
		fmt.Println("Request failed:", error)
		return
	}


	// defer means when main finishes running, close the stream
	defer response.Body.Close()

	// io will read the stream and store the data in body
	body, error := io.ReadAll(response.Body)


	// by default the API returns byte slices, convert it to string to read it
	var data map[string]interface{}

	error = json.Unmarshal(body, &data)

	fmt.Println(data["name"])

}