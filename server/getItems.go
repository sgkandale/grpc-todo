package server

import (
	"context"
	"log"

	item "grpc-todo/item"
	todoPB "grpc-todo/todoPB"
)

func (*Server) GetItems(ctx context.Context, r *todoPB.GetItemsRequest) (*todoPB.GetItemsResponse, error) {
	log.Println("GetItems() called")

	dbItems, err := item.GetAllItems()
	if err != nil {
		return nil, err
	}

	toRespItems := []*todoPB.Item{}

	for i := range dbItems {
		toRespItems = append(toRespItems, &todoPB.Item{
			Id:          dbItems[i].ID,
			Title:       dbItems[i].Title,
			Description: dbItems[i].Description,
			Closed:      dbItems[i].Closed,
		})
	}

	return &todoPB.GetItemsResponse{
		Items: toRespItems,
	}, nil

}
