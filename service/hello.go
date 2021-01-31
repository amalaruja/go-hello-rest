package service

import (
	"fmt"
    "net/http"
)

type Hello struct {
	Message string
}

func (hello *Hello) SayHello(w http.ResponseWriter, r *http.Request) {
	replyMessage := fmt.Sprintf("Hello from %s!\n", hello.Message)
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(replyMessage))
}