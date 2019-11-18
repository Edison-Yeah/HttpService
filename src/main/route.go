package main

import (
	"fmt"
	"net/http"
	"strings"
)

func searchHandler(w http.ResponseWriter, r *http.Request) {

	partsInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(partsInfo, "/")

	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	fmt.Print(action)

}
