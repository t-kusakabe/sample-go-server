package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sample(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("value: ", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello World.")
}

func main() {
	http.HandleFunc("/", sample)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenServe: ", err)
	}
}
