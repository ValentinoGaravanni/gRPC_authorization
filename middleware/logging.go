package middleware

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
)

func LoggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Printf("Request - Method : %s; Request: %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	if err != nil {
		st, _ := status.FromError(err)
		log.Printf("Error - Method:%s; Code:%s; Message:%s", info.FullMethod, st.Code(), st.Message())
	} else {
		log.Printf("Response - Method:%s; Response:%v", info.FullMethod, resp)
	}
	return resp, err
}
