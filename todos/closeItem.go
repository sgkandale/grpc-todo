package todos

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (todos *Todos) CloseItem(id string) error {

	for i := range todos.Items {
		if todos.Items[i].ID == id {
			todos.Items[i].Closed = true
			return nil
		}
	}

	return status.Error(codes.NotFound, "item not found")
}
