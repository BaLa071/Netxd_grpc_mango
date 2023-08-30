package main

import (
	"context"
	"fmt"
	"log"

	pb "DemoProject/Netxd_Customer_proto/Netxd_Customer"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewNetxd_DalServiceClient(conn)

	response, err := client.CreateCustomer(context.Background(), &pb.CustomerProfile{CustomerId: 1,FirstName: "balu"})
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}

	fmt.Printf("Response: %s\n", response.CustomerId)
}
