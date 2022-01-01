package item

type Item struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	Title       string `json:"title" bson:"title,omitempty"`
	Description string `json:"description" bson:"description,omitempty"`
	Closed      bool   `json:"closed" bson:"closed,omitempty"`
}
