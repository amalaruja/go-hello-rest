package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/amalaruja/go-hello-rest/service"
)

const apiPrefix = "/hello-rest-service/"

func main() {
	namePtr := flag.String("name", "DC 1", "Name")
	flag.Parse()

	helloService := service.Hello{Message: *namePtr}

	http.HandleFunc(fmt.Sprintf("%s%s", apiPrefix, "hello"), helloService.SayHello)
	http.ListenAndServe(":8080", nil)
}
