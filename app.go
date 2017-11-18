package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	initLine()

	router := httprouter.New()
	router.GET("/test", testHandler)
	http.ListenAndServe(":8080", router)
	fmt.Println("httprouter is on! やった！！")
}
