package server

import (
	"context"
	"log"

	item "grpc-todo/item"
	todoPB "grpc-todo/todoPB"
)

func (*Server) CreateItem(ctx context.Context, r *todoPB.Item) (*todoPB.Item, error) {
	log.Println("CreateItem() called")

	newItem, err := item.NewItem(
		r.GetTitle(),
		r.GetDescription(),
	)

	if err != nil {
		return nil, err
	}

	err = newItem.InsertInDB()
	if err != nil {
		return nil, err
	}

	return &todoPB.Item{
		Id:          newItem.ID,
		Title:       newItem.Title,
		Description: newItem.Description,
		Closed:      newItem.Closed,
	}, nil

}
