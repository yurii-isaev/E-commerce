package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomeRouterHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello word"))
}
