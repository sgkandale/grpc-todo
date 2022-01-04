package item

type Item struct {
	ID          string `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Closed      bool   `json:"closed" db:"closed"`
}
