package clientEngine

import (
	"context"
	"log"
	"os"
	"time"

	GrpcEngine "github.com/nusa-exchange/pkg/Grpc/engine"
	"google.golang.org/grpc"
)

type GrpcMatchingClient struct {
	connect *grpc.ClientConn
	client  GrpcEngine.MatchingEngineServiceClient
}

func NewMatchingClient() *GrpcMatchingClient {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(os.Getenv("MATCHING_ENGINE_URL"), grpc.WithInsecure(), grpc.WithBackoffMaxDelay(5*time.Second))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	return &GrpcMatchingClient{
		connect: conn,
		client:  GrpcEngine.NewMatchingEngineServiceClient(conn),
	}
}

func (c *GrpcMatchingClient) CalcMarketOrder(in *GrpcEngine.CalcMarketOrderRequest, opts ...grpc.CallOption) (*GrpcEngine.CalcMarketOrderResponse, error) {
	return c.client.CalcMarketOrder(context.Background(), in, opts...)
}

func (c *GrpcMatchingClient) FetchMarketPrice(in *GrpcEngine.FetchMarketPriceRequest, opts ...grpc.CallOption) (*GrpcEngine.FetchMarketPriceResponse, error) {
	return c.client.FetchMarketPrice(context.Background(), in, opts...)
}

func (c *GrpcMatchingClient) FetchOrderBook(in *GrpcEngine.FetchOrderBookRequest, opts ...grpc.CallOption) (*GrpcEngine.FetchOrderBookResponse, error) {
	return c.client.FetchOrderBook(context.Background(), in, opts...)
}

func (c *GrpcMatchingClient) FetchOrder(in *GrpcEngine.FetchOrderRequest, opts ...grpc.CallOption) (*GrpcEngine.FetchOrderResponse, error) {
	return c.client.FetchOrder(context.Background(), in, opts...)
}

func (c *GrpcMatchingClient) Close() error {
	return c.connect.Close()
}
