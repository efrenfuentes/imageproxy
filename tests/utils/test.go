package utils

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
)

func MakeRequest(method, urlStr string, body io.Reader, router *mux.Router) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, urlStr, body)

	if err != nil {
		log.Println(err)
	}

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	return w
}
