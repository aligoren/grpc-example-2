package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc_example/gen/protos"
	"io"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Error while connecting to gRPC server. Error: %v", err)
	}

	client := pb.NewHealthServiceClient(conn)

	resp, _ := client.GetSystemHealth(context.Background(), &pb.HealthRequest{
		Type: "disk",
	})

	stream, _ := client.GetCpuData(context.Background())

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			break
		}

		log.Printf("Response: %s", res.Data)
	}

	log.Printf("Disk info coming from server. OS: %s, data: %s", resp.Os, resp.Data)
}
