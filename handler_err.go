package main

import (
	"net/http"
)

func handlererr(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 400, "Something went wrong")
}
