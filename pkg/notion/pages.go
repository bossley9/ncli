package notion

import "time"

// https://developers.notion.com/reference/page
type Page struct {
	Object         string          `json:"object"` // always "page"
	ID             string          `json:"id"`
	CreatedTime    time.Time       `json:"created_time"`
	CreatedBy      PartialUser     `json:"created_by"`
	LastEditedTime time.Time       `json:"last_edited_time"`
	LastEditedBy   PartialUser     `json:"last_edited_by"`
	Archived       bool            `json:"archived"`
	Icon           Emoji           `json:"icon"`
	Cover          File            `json:"cover"`
	Properties     struct{}        `json:"properties"`
	Parent         WorkspaceParent `json:"parent"`
	URL            string          `json:"url"`
}
