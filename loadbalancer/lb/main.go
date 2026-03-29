package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("registering a load balancer")
	AddServer("5001")
	AddServer("5002")
	AddServer("5003")

	http.HandleFunc("/", LoadBalancer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("load balancer failed to start")
		return
	}
}
