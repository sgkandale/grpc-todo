package item

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewItemInstance(id string) (*Item, error) {

	if id == "" {
		return nil, status.Error(codes.InvalidArgument, "id cannot be empty")
	}
	if len(id) != 15 {
		return nil, status.Error(codes.InvalidArgument, "id must be 15 characters long")
	}

	return &Item{
		ID: id,
	}, nil

}
