package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"url-shortener/middleware"
	"url-shortener/storage"

	pb "url-shortener/proto"
)

const (
	BaseURL    = "http://localhost:8081/short/"
	GRPCPort   = ":8080"
	HTTPPort   = ":8081"
	DBFilePath = "urlshortener.db"
)

type server struct {
	pb.UnimplementedURLShortenerServer
	save *storage.Storage
}

func (s *server) URLShort(ctx context.Context, req *pb.ShortenRequest) (*pb.ShortenResponse, error) {
	if req.OriginalURL == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Original URL cannot be empty")
	}

	shortID := fmt.Sprintf("%d", time.Now().UnixNano())
	if err := s.save.AddURL(req.OriginalURL, shortID); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to shorten URL: %v", err)
	}

	return &pb.ShortenResponse{ShortURL: BaseURL + shortID}, nil
}

func (s *server) GetOriginalURL(ctx context.Context, req *pb.OriginalRequest) (*pb.OriginalResponse, error) {
	shortID := strings.TrimPrefix(req.ShortURL, BaseURL)
	original, err := s.save.GetOrginalURL(shortID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "Short URL not found")
		}
		return nil, status.Errorf(codes.Internal, "Failed to fetch original URL")
	}
	return &pb.OriginalResponse{OriginalURL: original}, nil
}

func (s *server) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortID := strings.TrimPrefix(r.URL.Path, "/short/")
	original, err := s.save.GetOrginalURL(shortID)
	if err != nil {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, original, http.StatusFound)
}

func main() {
	store := storage.NewStorage(DBFilePath)
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.LoggingInterceptor,
			middleware.RateLimitingInterceptor(middleware.NewRateLimiter(10, time.Second)),
		),
	)
	pb.RegisterURLShortenerServer(grpcServer, &server{save: store})

	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterURLShortenerHandlerFromEndpoint(ctx, mux, GRPCPort, opts); err != nil {
		log.Fatalf("Failed to register gRPC-Gateway: %v", err)
	}

	httpMux := http.NewServeMux()

	// ✅ Додаємо перевірку стану сервера
	httpMux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// ✅ Додаємо хендлер для кореневого `/`
	httpMux.Handle("/", http.FileServer(http.Dir("./web")))

	// ✅ Підключаємо gRPC-Gateway на `/api/v1/`
	httpMux.Handle("/api/v1/", http.StripPrefix("/api/v1", mux))

	// ✅ Додаємо обробник редіректу
	httpMux.HandleFunc("/short/", (&server{save: store}).RedirectHandler)

	// ✅ Додаємо можливість віддавати статичні файли, якщо `web/` існує
	httpMux.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("./web"))))

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}).Handler(httpMux)

	httpServer := &http.Server{Addr: HTTPPort, Handler: handler}
	log.Println("🚀 Starting HTTP server on", HTTPPort)
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	lis, err := net.Listen("tcp", GRPCPort)
	if err != nil {
		log.Fatalf("gRPC server error: %v", err)
	}
	log.Println("🚀 Starting gRPC server on", GRPCPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRPC server error: %v", err)
	}
}
