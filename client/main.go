package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"strings"
	"time"
	"url-shortener/auth"

	"google.golang.org/grpc"
	pb "url-shortener/url-shortener/pb"
)

func shortenURL(client pb.URLShortenerClient, ctx context.Context) {
	fmt.Print("Please enter URL which you would like to shorten: ")
	reader := bufio.NewReader(os.Stdin)
	originalURL, _ := reader.ReadString('\n')
	originalURL = strings.TrimSpace(originalURL)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.URLShort(ctx, &pb.ShortenRequest{OriginalURL: originalURL}, grpc.Header(&metadata.MD{}))
	if err != nil {
		log.Fatalf("Error calling ShortURL: %v", err)
	}
	fmt.Printf("Shortened URL: %s\n", resp.ShortURL)
}

func getOriginalURL(client pb.URLShortenerClient, ctx context.Context) {
	fmt.Print("Enter the short URL to receive the original from DB: ")
	reader := bufio.NewReader(os.Stdin)
	shortURL, _ := reader.ReadString('\n')
	shortURL = strings.TrimSpace(shortURL)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.GetOriginalURL(ctx, &pb.OriginalRequest{ShortURL: shortURL})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			log.Printf("Error: %s", st.Message())
		} else {
			log.Fatalf("Uknown error: %v", err)
		}
		return
	}
	fmt.Printf("Original URL: %s\n", resp.OriginalURL)
}

func getStats(client pb.URLShortenerClient, shortURL string, ctx context.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.GetStats(ctx, &pb.StatRequest{ShortURL: shortURL})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			log.Printf("Error: %s", st.Message())
		} else {
			log.Fatalf("Error calling GetStats: %v", err)
		}
		return
	}
	log.Printf("Statistics for %s - Click coutn: %d, Last Access: %s", shortURL, resp.ClickCount, resp.LastAccess)
}

func main() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewURLShortenerClient(conn)

	//Token generation
	token, err := auth.GenerateJWT(os.Getenv("SECRET_KEY"))
	if err != nil {
		log.Fatalf("Failed to generate token: %v", err)
	}
	//token in metadata
	md := metadata.Pairs("authorization", token)
	baseCtx := metadata.NewOutgoingContext(context.Background(), md)

	for {
		fmt.Println("\n=============SHORT LINK=============")
		fmt.Println("1. Shorten a URL")
		fmt.Println("2. Get original URL")
		fmt.Println("3. Statistics")
		fmt.Println("4. Exit")
		fmt.Println("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			shortenURL(client, baseCtx)
		case 2:
			getOriginalURL(client, baseCtx)
		case 3:
			fmt.Print("Enter the short URL to get statistics: ")
			reader := bufio.NewReader(os.Stdin)
			shortURL, _ := reader.ReadString('\n')
			shortURL = strings.TrimSpace(shortURL)

			getStats(client, shortURL, baseCtx)
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice.Please try again")
			//activation Jetbrains
		}
	}

}
