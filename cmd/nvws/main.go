package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var laddr = flag.String("laddr", ":80", "listen on address `LADDR`")

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello\n")
}

func main() {
	flag.Parse()
	http.HandleFunc("/", rootHandler)
	err := http.ListenAndServe(*laddr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
