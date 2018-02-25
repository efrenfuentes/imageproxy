package controllers

import (
	"net/http"
	"fmt"
)

func PingIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}
