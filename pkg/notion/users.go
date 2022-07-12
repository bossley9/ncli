package notion

type PartialUser struct {
	Object string `json:"object"` // always "user"
	ID     string `json:"id"`
}
