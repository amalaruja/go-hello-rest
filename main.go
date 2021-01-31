package main

import (
	"net/http"
	"flag"

	"github.com/amalaruja/go-hello-rest/service"
)

func main() {
	namePtr := flag.String("name", "DC 1", "Name")
	flag.Parse()

	helloService := service.Hello{Message: *namePtr}

    http.HandleFunc("/hello", helloService.SayHello)
    http.ListenAndServe(":8080", nil)
}