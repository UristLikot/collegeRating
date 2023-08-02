package main

import "net/http"

func main() {
	telegramBot()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("200"))
}
