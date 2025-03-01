package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"

	"net/http"
	"time"
	"url-shortener/middleware"
	"url-shortener/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "url-shortener/proto" /// Import generated files
)

type server struct {
	pb.UnimplementedURLShortenerServer
	save *storage.Storage
}

func (s *server) URLShort(ctx context.Context, req *pb.ShortenRequest) (*pb.ShortenResponse, error) {
	short := "https://short.link/" + fmt.Sprintf("%d", time.Now().UnixNano())

	//Add to DB
	if err := s.save.AddURL(req.OriginalURL, short); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to shorten URL: %v", err)
	}
	return &pb.ShortenResponse{ShortURL: short}, nil
}

func (s *server) GetOriginalURL(ctx context.Context, req *pb.OriginalRequest) (*pb.OriginalResponse, error) {
	original, err := s.save.GetOrginalURL(req.ShortURL)
	if err != nil {
		//Check if url exists in DB
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "Short URL not found %v", req.ShortURL)
		}
		return nil, status.Errorf(codes.Internal, "Failed to fetch original URL: %v", err)
	}
	if err := s.save.ClickCountInc(req.ShortURL); err != nil {
		log.Printf("Failed to increment click count for %s:  %v", req.ShortURL, err)
	}
	return &pb.OriginalResponse{OriginalURL: original}, nil
}

func (s *server) GetStats(ctx context.Context, req *pb.StatRequest) (*pb.StatResponse, error) {
	//Receiving stats
	clickCount, err := s.save.GetStats(req.ShortURL)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Short URL not found: %v", err)
	}

	//Reply
	return &pb.StatResponse{
		ClickCount: int64(int32(clickCount)),
		LastAccess: time.Now().Format(time.RFC3339), //Last time format
	}, nil
}

func main() {

	// DB init
	store := storage.NewStorage("urlshortener.db")

	// gRPC server
	rl := middleware.NewRateLimiter(10, time.Second)
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.LoggingInterceptor,
			middleware.RateLimitingInterceptor(rl),
		),
	)
	pb.RegisterURLShortenerServer(grpcServer, &server{save: store})

	// gRPC-Gateway mux
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())} // заміна WithInsecure()
	err := pb.RegisterURLShortenerHandlerFromEndpoint(ctx, mux, ":8080", opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC-Gateway: %v", err)
	}

	// HTTP server for gRPC-Gateway
	httpMux := http.NewServeMux()
	httpMux.Handle("/api/", http.StripPrefix("/api", mux))

	fs := http.FileServer(http.Dir("./web"))
	httpMux.Handle("/", fs)

	// CORS adding
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, //all domains
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})
	handler := c.Handler(httpMux)

	// HTTP CORS launch
	srv := &http.Server{
		Addr:    ":8081",
		Handler: handler,
	}

	log.Println("Starting gRPC-Gateway on :8081")
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	// gRPC server launch
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen on :8080: %v", err)
	}
	log.Println("Starting gRPC server on :8080")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
