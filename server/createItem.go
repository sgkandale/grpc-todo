package server

import (
	"context"
	"log"

	"todo/item"
	"todo/todoPB"
	"todo/todos"
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

	todos.TodoList.AddItem(newItem)

	return &todoPB.Item{
		Id:          newItem.ID,
		Title:       newItem.Title,
		Description: newItem.Description,
		Closed:      newItem.Closed,
	}, nil

}
