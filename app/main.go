package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func newHandler(name string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Hello, App.\n")

		w.Header().Add("content-type", "text/html")
		w.WriteHeader(200)

		fmt.Fprintf(w, "<html><head><meta charset=\"utf8\"/><title>%[1]sの部屋</title></head><body><h1>%[1]sの部屋</h1><p>ようこそ！あなたは、", name)
		fmt.Fprintf(w, "undefined")
		fmt.Fprintf(w, "人目のお客さまです。</p></body></html>")
	};
}

func main() {
	port, found := os.LookupEnv("PORT")
	if !found {
		port = "80"
	}

	addr := fmt.Sprintf(":%s", port)

	name, found := os.LookupEnv("NAME")
	if !found {
		name = "undefined"
	}

	handler := newHandler(name)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(addr, nil))
}
