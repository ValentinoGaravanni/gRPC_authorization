package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "url-shortener/url-shortener/pb" /// Import generated files
)

type server struct {
	pb.UnimplementedURLShortenerServer
}

func (s *server) URLShort(ctx context.Context, req *pb.ShortenRequest) (*pb.ShortenResponse, error) {
	short := "short.link/" + "948234"
	return &pb.ShortenResponse{ShortURL: short}, nil
}

func (s *server) GetOriginalURL(ctx context.Context, req *pb.OriginalRequest) (*pb.OriginalResponse, error) {
	original := "https://www.google.com"
	return &pb.OriginalResponse{OriginalURL: original}, nil
}

func (s *server) GetStat(ctx context.Context, req *pb.StatRequest) (*pb.StatResponse, error) {
	stats := &pb.StatResponse{
		ClickCount: 42,
		LastAccess: "2025-01-16T12:00:00Z",
	}
	return stats, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: &v", err)
	}

	grpcServer := grpc.NewServer() //Server created
	pb.RegisterURLShortenerServer(grpcServer, &server{})

	log.Println("Server is running on port :50052 ") //Server launch
	if err := grpcServer.Serve(listener); err != nil {
	}
	log.Fatalf("Failed to serve: %v", err)

}
