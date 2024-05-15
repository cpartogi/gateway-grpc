package helper

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

func GenerateRequestID() (ctx context.Context, requestID string) {

	requestID = uuid.New().String()

	md := metadata.Pairs("requestid", requestID)

	return metadata.NewOutgoingContext(context.Background(), md), requestID
}
