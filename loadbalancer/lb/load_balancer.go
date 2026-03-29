package main

import (
	"io"
	"log"
	"net/http"
	"sync/atomic"
)

func LoadBalancer(w http.ResponseWriter, r *http.Request) {
	log.Println("enter load balancer")
	bestServer := GetBestServer()
	if bestServer == nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	// atomic ensures that the server is locked until sometime so that the request of person A completes hence avoiding race condition
	atomic.AddUint64(&bestServer.RequestsCount, 1)
	// no matter whether request is success or not just decrement the requestCount
	defer atomic.AddUint64(&bestServer.RequestsCount, ^uint64(0))
	log.Printf("received the request , port forwarding to %s", bestServer.Url)
	resp, err := http.Get("http://localhost:" + bestServer.Url)
	if err != nil {
		log.Println("error in load balancer")
		return
	}
	defer resp.Body.Close()

	for header, headerValue := range resp.Header {
		for _, value := range headerValue {
			w.Header().Add(header, value)
		}
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
	log.Println("end load balancer")
}
