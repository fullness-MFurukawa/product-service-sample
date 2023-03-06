package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
ハンドラ
*/
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Product Service!!")
}

/*
エントリーポイント
*/
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
