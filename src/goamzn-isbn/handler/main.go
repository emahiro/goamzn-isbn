package handler

import "net/http"

func Top(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome goamzn!"))

}
