package handler

import (
	"fmt"
	"net/http"
)

type ShowHelloWorldHandler struct {
}

func NewHelloWorldHandler() *ShowHelloWorldHandler {
	return &ShowHelloWorldHandler{}
}

func (h *ShowHelloWorldHandler) ShowHelloWorld(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "Hello World\n")
}
