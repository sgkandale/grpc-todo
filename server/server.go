package server

import (
	todoPB "grpc-todo/todoPB"
)

//Server exposed
type Server struct {
	todoPB.UnimplementedTodoServiceServer
}
