package server

import (
	"context"
	"log"

	"todo/todoPB"
	"todo/todos"
)

func (*Server) DeleteItem(ctx context.Context, r *todoPB.Item) (*todoPB.GeneralResponse, error) {
	log.Println("DeleteItem() called")

	err := todos.TodoList.DeleteItem(r.GetId())
	if err != nil {
		return nil, err
	}

	return &todoPB.GeneralResponse{
		Message: "Item deleted",
	}, nil

}
