package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func newHandler(name, endpoint string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(endpoint)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "%s", err)
			return
		}
		defer resp.Body.Close()

		w.Header().Add("content-type", "text/html")
		w.WriteHeader(200)

		fmt.Fprintf(w, "<html><head><meta charset=\"utf8\"/><title>%[1]sの部屋</title></head><body><h1>%[1]sの部屋</h1><p>ようこそ！あなたは、", name)
		if _, err := io.Copy(w, resp.Body); err != nil {
			log.Printf("%s", err)
		}
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

	endpoint, found := os.LookupEnv("EXTERNAL_ENDPOINT")
	if !found {
		log.Fatal("missing `EXTERNAL_ENDPOINT`")
	}

	handler := newHandler(name, endpoint)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(addr, nil))
}
