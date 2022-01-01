package todos

import (
	"todo/item"
)

type Todos struct {
	Items []item.Item `json:"items"`
}
