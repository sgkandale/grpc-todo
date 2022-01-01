package server

import (
	"context"
	"log"

	"todo/todoPB"
	"todo/todos"
)

func (*Server) CloseItem(ctx context.Context, r *todoPB.Item) (*todoPB.GeneralResponse, error) {
	log.Println("CloseItem() called")

	err := todos.TodoList.CloseItem(r.GetId())
	if err != nil {
		return nil, err
	}

	return &todoPB.GeneralResponse{
		Message: "Item closed",
	}, nil

}
