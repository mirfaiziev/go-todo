package handler

import (
	"fmt"
	"net/http"
)

func ShowHelloWorld(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "Hello World\n")
}
