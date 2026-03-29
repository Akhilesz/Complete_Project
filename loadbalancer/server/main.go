package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please provide port")
		return
	}
	port := os.Args[1]
	fmt.Printf("hello from port %s\n", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprint(w, "Hello world from", port)
		log.Println("number of bytes written are ", n, err)
		if err != nil {
			log.Fatal(err)
		}
	})
	fmt.Println("listening on port " + port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
