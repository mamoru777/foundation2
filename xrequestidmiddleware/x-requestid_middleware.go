package xrequestidmiddleware

import (
	"context"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
)

type ReqInterceptor struct{}

func NewReqInterceptor() *ReqInterceptor {
	return &ReqInterceptor{}
}

func (i *ReqInterceptor) RequestIDInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Println("Использована прослойка request id")
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if ok {
		//return nil, status.Errorf(codes.Internal, "metadata is not provided")
		if len(md.HeaderMD.Get("x_request_id")) == 0 {
			// X-Request-Id отсутствует, создаем новый и добавляем его в метаданные
			requestID := uuid.New().String()
			md.HeaderMD.Set("x_request_id", requestID)
		}
		// Продолжаем обработку запроса
		return handler(ctx, req)
	} else {
		requestID := uuid.New().String()
		ctx = context.WithValue(ctx, "x_request_id", requestID)
		return handler(ctx, req)
	}
}
