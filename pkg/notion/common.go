package notion

type Color string // enum

// https://developers.notion.com/reference/rich-text
type RichText struct {
	PlainText   string     `json:"plain_text"`
	Href        *string    `json:"href,omitempty"`
	Annotations Annotation `json:"annotations"`
	Type        string     `json:"type"` // one of "text", "mention", "equation"
	// TODO remove
	Text struct {
		Content string      `json:"content"`
		Link    interface{} `json:"link"`
	} `json:"text"`
}

// https://developers.notion.com/reference/rich-text#annotations
type Annotation struct {
	Bold          bool  `json:"bold"`
	Italic        bool  `json:"italic"`
	Strikethrough bool  `json:"strikethrough"`
	Underline     bool  `json:"underline"`
	Code          bool  `json:"code"`
	Color         Color `json:"color"`
}
