package server

import (
	"context"
	"gogateway/server/customerpb"
	"log"
)

type Server struct {
	customerpb.UnimplementedCustomerServer
}

func (Server) GetCustomer(ctx context.Context, req *customerpb.GetCustomersRequest) (*customerpb.GetCustomerResponse, error) {
	log.Println("processing a request...")
	return &customerpb.GetCustomerResponse{
		Customers: []string{"A310", "K423"},
	}, nil
}
