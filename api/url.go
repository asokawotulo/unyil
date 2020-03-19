package api

import (
	"net/http"
)

// Index yes
func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"hello": "world"}`))
}
