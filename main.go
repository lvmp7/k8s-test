package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, greeting("Code.education Rocks!"))
	})
	fmt.Println(http.ListenAndServe(":8000", nil))
}

func greeting(s string) string {
	return "<b>" + s + "</b>"
}
