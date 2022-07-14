package notion

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"git.sr.ht/~bossley9/sn/pkg/api"
)

// https://developers.notion.com/reference/block
type Block struct {
	Object         string      `json:"object"` // always "block"
	ID             string      `json:"id"`
	Parent         PageParent  `json:"parent"`
	Type           string      `json:"type"`
	CreatedTime    time.Time   `json:"created_time"`
	CreatedBy      PartialUser `json:"created_by"`
	LastEditedTime time.Time   `json:"last_edited_time"`
	LastEditedBy   PartialUser `json:"last_edited_by"`
	Archived       bool        `json:"archived"`
	HasChildren    bool        `json:"has_children"`

	Paragraph        *ParagraphBlock        `json:"paragraph,omitempty"`
	HeadingOne       *HeadingOneBlock       `json:"heading_1,omitempty"`
	HeadingTwo       *HeadingTwoBlock       `json:"heading_2,omitempty"`
	HeadingThree     *HeadingThreeBlock     `json:"heading_3,omitempty"`
	Quote            *QuoteBlock            `json:"quote,omitempty"`
	BulletedListItem *BulletedListItemBlock `json:"bulleted_list_item,omitempty"`
	NumberedListItem *NumberedListItemBlock `json:"numbered_list_item,omitempty"`
	Todo             *TodoBlock             `json:"to_do,omitempty"`
	Code             *CodeBlock             `json:"code,omitempty"`
}

func (block *Block) ToMarkdown(s *strings.Builder) {
	switch block.Type {
	case "paragraph":
		block.Paragraph.ToMarkdown(s)
	case "heading_1":
		block.HeadingOne.ToMarkdown(s)
	case "heading_2":
		block.HeadingTwo.ToMarkdown(s)
	case "heading_3":
		block.HeadingThree.ToMarkdown(s)
	case "quote":
		block.Quote.ToMarkdown(s)
	case "bulleted_list_item":
		block.BulletedListItem.ToMarkdown(s)
	case "numbered_list_item":
		block.NumberedListItem.ToMarkdown(s)
	case "to_do":
		block.Todo.ToMarkdown(s)
	case "code":
		block.Code.ToMarkdown(s)
	default:
		s.WriteString(fmt.Sprintf("UNIMPLEMENTED BLOCK TYPE: %s\n", block.Type))
	}
	s.WriteString("\n")
}

// https://developers.notion.com/reference/block#paragraph-blocks
type ParagraphBlock struct {
	RichText []RichText `json:"rich_text"`
	Color    Color      `json:"color"`
	Children []Block    `json:"children"`
}

func (p *ParagraphBlock) ToMarkdown(s *strings.Builder) {
	for _, rich := range p.RichText {
		rich.ToMarkdown(s)
	}
}

// https://developers.notion.com/reference/block#heading-one-blocks
type HeadingOneBlock struct {
	RichText []RichText `json:"rich_text"`
	Color    Color      `json:"color"`
}

func (h *HeadingOneBlock) ToMarkdown(s *strings.Builder) {
	s.WriteString("# ")
	for _, rich := range h.RichText {
		rich.ToMarkdown(s)
	}
}

// https://developers.notion.com/reference/block#heading-two-blocks
type HeadingTwoBlock struct {
	RichText []RichText `json:"rich_text"`
	Color    Color      `json:"color"`
}

func (h *HeadingTwoBlock) ToMarkdown(s *strings.Builder) {
	s.WriteString("## ")
	for _, rich := range h.RichText {
		rich.ToMarkdown(s)
	}
}

// https://developers.notion.com/reference/block#heading-three-blocks
type HeadingThreeBlock struct {
	RichText []RichText `json:"rich_text"`
	Color    Color      `json:"color"`
}

func (h *HeadingThreeBlock) ToMarkdown(s *strings.Builder) {
	s.WriteString("### ")
	for _, rich := range h.RichText {
		rich.ToMarkdown(s)
	}
}

// https://developers.notion.com/reference/block#quote-blocks
type QuoteBlock struct {
	RichText []RichText `json:"rich_text"`
	Color    Color      `json:"color"`
	Children []Block    `json:"children"`
}

func (q *QuoteBlock) ToMarkdown(s *strings.Builder) {
	s.WriteString("> ")
	for _, rich := range q.RichText {
		rich.ToMarkdown(s)
	}
}

// https://developers.notion.com/reference/block#bulleted-list-item-blocks
type BulletedListItemBlock struct {
	RichText []RichText `json:"rich_text"`
	Color    Color      `json:"color"`
	Children []Block    `json:"children"`
}

func (li *BulletedListItemBlock) ToMarkdown(s *strings.Builder) {
	s.WriteString("* ")
	for _, rich := range li.RichText {
		rich.ToMarkdown(s)
	}
}

// https://developers.notion.com/reference/block#numbered-list-item-blocks
type NumberedListItemBlock struct {
	RichText []RichText `json:"rich_text"`
	Color    Color      `json:"color"`
	Children []Block    `json:"children"`
}

func (li *NumberedListItemBlock) ToMarkdown(s *strings.Builder) {
	// need to fix block numbering - Notion's API doesn't play well with referencing consecutive blocks
	s.WriteString("1. ")
	for _, rich := range li.RichText {
		rich.ToMarkdown(s)
	}
}

// https://developers.notion.com/reference/block#to-do-blocks
type TodoBlock struct {
	RichText []RichText `json:"rich_text"`
	Checked  *bool      `json:"checked,omitempty"`
	Color    string     `json:"color"`
	Children []Block    `json:"children"`
}

func (t *TodoBlock) ToMarkdown(s *strings.Builder) {
	var state string
	if *t.Checked {
		state = "x"
	} else {
		state = " "
	}

	s.WriteString(fmt.Sprintf("- [%s] ", state))
	for _, rich := range t.RichText {
		rich.ToMarkdown(s)
	}
}

// https://developers.notion.com/reference/block#code-blocks
type CodeBlock struct {
	RichText []RichText `json:"rich_text"`
	Caption  []RichText `json:"caption"`
	Language string     `json:"language"` // enum
}

func (c *CodeBlock) ToMarkdown(s *strings.Builder) {
	s.WriteString(fmt.Sprintf("```%s\n", c.Language))
	for _, rich := range c.RichText {
		rich.ToMarkdown(s)
	}
	s.WriteString("```\n")
}

type RetrieveBlockChildrenResponse struct {
	Object     string   `json:"object"`
	Results    []Block  `json:"results"`
	NextCursor *string  `json:"next_cursor"`
	HasMore    bool     `json:"has_more"`
	Type       string   `json:"type"`
	Block      struct{} `json:"block"`
}

func RetrieveBlockChildren(blockID string, startCursor string, pageSize int) (*RetrieveBlockChildrenResponse, error) {
	qs := url.Values{}
	if len(startCursor) > 0 {
		qs.Set("start_cursor", startCursor)
	}
	if pageSize > 0 && pageSize <= 100 {
		qs.Set("page_size", strconv.Itoa(pageSize))
	}

	url := fmt.Sprintf("%s/blocks/%s/children?%s", NOTION_API_URL, blockID, qs.Encode())
	params := map[string]any{}
	headers := getHeaders()

	body, err := api.Fetch(url, "GET", params, headers)
	if err != nil {
		return nil, err
	}

	var response RetrieveBlockChildrenResponse
	json.Unmarshal(body, &response)

	return &response, nil
}
