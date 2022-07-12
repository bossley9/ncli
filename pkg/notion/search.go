package notion

type SearchResponse struct {
	Object         string   `json:"object"`
	Results        []Page   `json:"results"`
	NextCursor     *string  `json:"next_cursor"`
	HasMore        bool     `json:"has_more"`
	Type           string   `json:"type"`
	PageOrDatabase struct{} `json:"page_or_database"`
}
