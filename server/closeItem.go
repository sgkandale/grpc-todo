package server

import (
	"context"
	"log"

	item "grpc-todo/item"
	todoPB "grpc-todo/todoPB"
)

func (*Server) CloseItem(ctx context.Context, r *todoPB.Item) (*todoPB.GeneralResponse, error) {
	log.Println("CloseItem() called")

	itemInstance, err := item.NewItemInstance(r.GetId())
	if err != nil {
		return nil, err
	}

	err = itemInstance.LoadDetails()
	if err != nil {
		return nil, err
	}

	err = itemInstance.SetClosed()
	if err != nil {
		return nil, err
	}

	return &todoPB.GeneralResponse{
		Message: "Item closed",
	}, nil

}
