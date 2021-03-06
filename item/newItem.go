package item

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewItem(title string, description string) (Item, error) {

	blankItem := Item{}

	// title input checks

	if title == "" {
		return blankItem, status.Error(codes.InvalidArgument, "title cannot be empty")
	}
	if len(title) > 100 {
		return blankItem, status.Error(codes.InvalidArgument, "title cannot be longer than 100 characters")
	}

	// description input checks

	if description == "" {
		return blankItem, status.Error(codes.InvalidArgument, "description cannot be empty")
	}
	if len(description) > 500 {
		return blankItem, status.Error(codes.InvalidArgument, "description cannot be longer than 500 characters")
	}

	// return new item

	return Item{
		ID:          GererateId(),
		Title:       title,
		Description: description,
		Closed:      false,
	}, nil
}
