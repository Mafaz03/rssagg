package main

import "net/http"

func handler_err(w http.ResponseWriter, r *http.Request) {
	errResponse(w, 500, "something went wrong")
}
