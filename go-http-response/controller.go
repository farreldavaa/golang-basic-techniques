package main

import (
	"encoding/json"
	"net/http"
)

type Data struct {
	Nama   string
	Nim    string
	Alamat string
}

var data = []Data{}

func Students(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		var newData = Data{}
		err := json.NewDecoder(req.Body).Decode(&newData)

		res.Header().Set("Content-Type", "application.json")
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(res).Encode(map[string]string{"error": "Bad Request"})
			return
		}

		data = append(data, newData)

		res.WriteHeader(http.StatusCreated)
		return
	}

	HandleNotAllowed(res, req)
}

func AllData(res http.ResponseWriter, req *http.Request){
	if req.Method == "GET" {
		res.Header().Set("Content-Type", "application.json")
		res.WriteHeader(http.StatusOK)
		D, _ := json.Marshal(data)
		res.Write(D)
		return
	}

	HandleNotAllowed(res, req)
}

func HandleNotAllowed(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(res).Encode(map[string]string{"error": "Method Not Allowed"})
}