package main

import (
	"context"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"golang.org/x/sys/windows"
	"google.golang.org/grpc"
	pb "grpc_example/gen/protos"
	"log"
	"net"
	"runtime"
	"time"
)

type healthServer struct {
	pb.UnimplementedHealthServiceServer
}

func diskUsage(path string) string {

	pathPtr, _ := windows.UTF16PtrFromString(path)

	var free, total, availableSize uint64

	windows.GetDiskFreeSpaceEx(pathPtr, &free, &total, &availableSize)

	return fmt.Sprintf("Free: %d, Total: %d, Available: %d", free, total, availableSize)
}

func (h *healthServer) GetSystemHealth(ctx context.Context, request *pb.HealthRequest) (*pb.HealthResponse, error) {

	var data string

	if request.Type == "disk" {
		data = diskUsage("c:\\")
	}

	return &pb.HealthResponse{
		Os:   runtime.GOOS,
		Data: data,
	}, nil
}

func (h *healthServer) GetCpuData(stream pb.HealthService_GetCpuDataServer) error {

	var count int = 0

	for {

		percent, _ := cpu.Percent(time.Second, false)

		stream.Send(&pb.CpuResponse{
			Data: fmt.Sprintf("CPU Usage: %v\n", percent),
		})

		count = count + 1

		if count == 5 {
			break
		}
	}

	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("Error while listening port :8080. Error: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterHealthServiceServer(grpcServer, &healthServer{})

	log.Println("gRPC server is running on :8080")

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Error while serving grpcServer. Error: %v", err)
	}

}
