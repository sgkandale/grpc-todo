package server

import (
	"todo/todoPB"
)

//Server exposed
type Server struct {
	todoPB.UnimplementedTodoServiceServer
}
