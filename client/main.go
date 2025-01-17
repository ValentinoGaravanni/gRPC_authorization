package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "url-shortener/url-shortener/pb"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewURLShortenerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Shortener Test
	shortenResp, err := client.URLShort(ctx, &pb.ShortenRequest{
		OriginalURL: "https://example.com",
	})
	if err != nil {
		log.Fatalf("Error calling ShortenURL: %v", err)
	}
	log.Printf("Shortened URL: %s", shortenResp.ShortURL)
}
