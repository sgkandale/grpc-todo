package server

import (
	"context"
	"log"

	"todo/todoPB"
	"todo/todos"
)

func (*Server) GetItems(ctx context.Context, r *todoPB.GetItemsRequest) (*todoPB.GetItemsResponse, error) {
	log.Println("GetItems() called")

	insertedItems := todos.TodoList.Items

	toRespItems := []*todoPB.Item{}

	for i := range insertedItems {
		toRespItems = append(toRespItems, &todoPB.Item{
			Id:          insertedItems[i].ID,
			Title:       insertedItems[i].Title,
			Description: insertedItems[i].Description,
			Closed:      insertedItems[i].Closed,
		})
	}

	return &todoPB.GetItemsResponse{
		Items: toRespItems,
	}, nil

}
