package main

import "net/http"

func main() {

	http.HandleFunc("/search/", searchHandler)
	http.ListenAndServe("8000", nil)
}
