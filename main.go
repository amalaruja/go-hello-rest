package main

import (
	"net/http"
	
	"github.com/amalaruja/go-hello-rest/service"
)

func main() {
	helloService := service.Hello{Message: "DC 1"}
    http.HandleFunc("/hello", helloService.SayHello)
    http.ListenAndServe(":8080", nil)
}
