package notion

import (
	"fmt"
	"time"
)

type Color string // enum

// https://developers.notion.com/reference/rich-text
type RichText struct {
	PlainText   string     `json:"plain_text"`
	Href        *string    `json:"href,omitempty"`
	Annotations Annotation `json:"annotations"`
	Type        string     `json:"type"` // one of "text", "mention", "equation"
}

func (rich *RichText) ToMarkdown() string {
	delimiter := ""

	if rich.Annotations.Bold {
		delimiter = "**"
	} else if rich.Annotations.Italic {
		delimiter = "_"
	} else if rich.Annotations.Strikethrough {
		delimiter = "~~"
	} else if rich.Annotations.Code {
		delimiter = "`"
	}

	return fmt.Sprintf("%s%s%s", delimiter, rich.PlainText, delimiter)
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

// https://developers.notion.com/reference/file-object
type File struct {
	Type string `json:"type"`

	External struct {
		URL string `json:"url"`
	} `json:"external"`
	File struct {
		URL        string    `json:"url"`
		ExpiryTime time.Time `json:"expiry_time"`
	} `json:"file"`
}

// https://developers.notion.com/reference/emoji-object
type Emoji struct {
	Type  string `json:"type"`
	Emoji string `json:"emoji"`
}

// https://developers.notion.com/reference/parent-object#page-parent
type PageParent struct {
	Type   string `json:"type"` // always "page_id"
	PageID string `json:"page_id"`
}

// https://developers.notion.com/reference/parent-object#workspace-parent
type WorkspaceParent struct {
	Type      string `json:"type"`      // always "workspace"
	Workspace bool   `json:"workspace"` // always true
}

// https://developers.notion.com/reference/parent-object#block-parent
type BlockParent struct {
	Type    string `json:"type"` // always "block_id"
	BlockID string `json:"block_id"`
}
