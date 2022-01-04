package item

type Item struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Closed      bool   `json:"closed"`
}
