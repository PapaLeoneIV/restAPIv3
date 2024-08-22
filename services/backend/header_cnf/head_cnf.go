package header_cnf

import (
	"net/http"
	"fmt"
)

func SetHeaders(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Printf("Header set\n")
		next(w, r)
	}
}