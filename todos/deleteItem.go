package todos

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (todos *Todos) DeleteItem(id string) error {
	for i, item := range todos.Items {
		if item.ID == id {
			todos.Items = append(
				(todos.Items)[:i],
				(todos.Items)[i+1:]...,
			)
			return nil
		}
	}

	return status.Error(codes.NotFound, "item not found")

}
