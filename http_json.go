package main

import (
	"encoding/json"
	"io"
	"net/http"
)

var count = 0

func hello(w http.ResponseWriter, r *http.Request) {
	println(r.Host)
	io.WriteString(w, "fuckyou")
}
func chat(w http.ResponseWriter, r *http.Request) {
	count++
	println(count)
	cb := r.FormValue("callback")
	// println(cb)
	f := map[string]interface{}{
		"name": "fuck",
		"age":  6,
	}
	str, _ := json.Marshal(f)
	if cb != "" {
		w.Write([]byte(cb))
		w.Write([]byte("("))
	}
	w.Write(str)
	if cb != "" {
		w.Write([]byte(")"))
	}
}
func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/chat", chat)
	http.ListenAndServe(":9090", nil)
}
