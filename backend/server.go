package main

import (
	"fmt"
	"log"
	"main/addition"
	"main/multiplication"
	"main/proto"
	"net"

	"google.golang.org/grpc"
)

func main() {
	go func() {
		listener, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatal(err)
		}
		s := grpc.NewServer()
		proto.RegisterAdditionServiceServer(s, &addition.Server{})
		log.Println("Addition gRPC Server listening on :50051")
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
		// 2. Call the Multiplication Client
		fmt.Println("Requesting Multiplication: 10 * 4...")
		result, err := multiplication.MultiplyViaGRPC(10, 4)
		fmt.Printf("🎯 Final Result from Microservices: %d\n", result)
	}()
	var a, b int64
	fmt.Println("------------------------------------")
	fmt.Print("Enter first number to multiply: ")
	fmt.Scanln(&a)
	fmt.Print("Enter second number to multiply: ")
	fmt.Scanln(&b)
	fmt.Println("------------------------------------")

	// 3. Call the Multiplication Client using the inputs
	fmt.Printf("🚀 Requesting Multiplication: %d * %d via gRPC...\n", a, b)

	// Note: I removed the second return 'err' if your MultiplyViaGRPC doesn't return one yet
	result, err := multiplication.MultiplyViaGRPC(a, b)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("🎯 Final Result: %d\n", result)
}
