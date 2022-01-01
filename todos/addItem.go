package todos

import "todo/item"

func (todos *Todos) AddItem(newItem item.Item) {
	todos.Items = append(
		todos.Items,
		newItem,
	)
}
