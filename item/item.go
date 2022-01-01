package item

type Item struct {
	ID          string `json:"id" bson:"_id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Closed      bool   `json:"closed" bson:"closed"`
}
