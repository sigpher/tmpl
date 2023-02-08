package main

import (
	"fmt"
	"net/http"
	"tmpl/handler"
)

func main() {
	http.HandleFunc("/", handler.HandleInfo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP Server failed, err:", err)
		return
	}
}
