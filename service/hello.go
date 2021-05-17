package service

import (
	"fmt"
	"net/http"
)

type Hello struct {
	Message string
}

func (hello *Hello) SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received with headers:\n", r.Header)
	replyMessage := fmt.Sprintf("Hello from %s!\n", hello.Message)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(replyMessage))
}
