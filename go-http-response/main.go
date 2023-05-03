package main

import (
	"fmt"
	"net/http"
)

func main() {

	var PORT = "5050"
	mux := http.NewServeMux()

	mux.HandleFunc("/nama", Students)
	mux.HandleFunc("/semuadata", AllData)
	fmt.Println("Server running on port " + PORT)

	err := http.ListenAndServe(":" + PORT, mux)

	if err != nil {
		fmt.Println(err)
		return
	}

}
