package notion

import (
	"fmt"
	"strings"
	"time"
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

func (block *Block) ToMarkdown() string {
	switch block.Type {
	case "paragraph":
		return block.Paragraph.ToMarkdown()
	case "heading_1":
		return block.HeadingOne.ToMarkdown()
	case "heading_2":
		return block.HeadingTwo.ToMarkdown()
	case "heading_3":
		return block.HeadingThree.ToMarkdown()
	case "quote":
		return block.Quote.ToMarkdown()
	case "bulleted_list_item":
		return block.BulletedListItem.ToMarkdown()
	case "numbered_list_item":
		return block.NumberedListItem.ToMarkdown()
	case "to_do":
		return block.Todo.ToMarkdown()
	case "code":
		return block.Code.ToMarkdown()
	default:
		return fmt.Sprintf("UNIMPLEMENTED BLOCK TYPE: %s\n", block.Type)
	}
}

// https://developers.notion.com/reference/block#paragraph-blocks
type ParagraphBlock struct {
	RichText []RichText `json:"rich_text"`
	Color    Color      `json:"color"`
	// TODO children blocks
	Children []Block `json:"children"`
}

func (paragraph *ParagraphBlock) ToMarkdown() string {
	var output strings.Builder
	for _, rich := range paragraph.RichText {
		output.WriteString(rich.ToMarkdown())
	}
	return output.String()
}

// https://developers.notion.com/reference/block#heading-one-blocks
type HeadingOneBlock struct {
	RichText []RichText `json:"rich_text"`
	Color    Color      `json:"color"`
}

func (heading *HeadingOneBlock) ToMarkdown() string {
	var output strings.Builder
	for _, rich := range heading.RichText {
		output.WriteString(fmt.Sprintf("# %s", rich.PlainText))
	}
	return output.String()
}

// https://developers.notion.com/reference/block#heading-two-blocks
type HeadingTwoBlock struct {
	RichText []RichText `json:"rich_text"`
	Color    Color      `json:"color"`
}

func (heading *HeadingTwoBlock) ToMarkdown() string {
	var output strings.Builder
	for _, rich := range heading.RichText {
		output.WriteString(fmt.Sprintf("## %s", rich.PlainText))
	}
	return output.String()
}

// https://developers.notion.com/reference/block#heading-three-blocks
type HeadingThreeBlock struct {
	RichText []RichText `json:"rich_text"`
	Color    Color      `json:"color"`
}

func (heading *HeadingThreeBlock) ToMarkdown() string {
	var output strings.Builder
	for _, rich := range heading.RichText {
		output.WriteString(fmt.Sprintf("### %s", rich.PlainText))
	}
	return output.String()
}

// https://developers.notion.com/reference/block#quote-blocks
type QuoteBlock struct {
	RichText []RichText `json:"rich_text"`
	Color    Color      `json:"color"`
	Children []Block    `json:"children"`
}

func (quote *QuoteBlock) ToMarkdown() string {
	var output strings.Builder
	for _, rich := range quote.RichText {
		output.WriteString(fmt.Sprintf("> %s", rich.PlainText))
	}
	return output.String()
}

// https://developers.notion.com/reference/block#bulleted-list-item-blocks
type BulletedListItemBlock struct {
	RichText []RichText `json:"rich_text"`
	Color    Color      `json:"color"`
	Children []Block    `json:"children"`
}

func (listItem *BulletedListItemBlock) ToMarkdown() string {
	var output strings.Builder
	for _, rich := range listItem.RichText {
		output.WriteString(fmt.Sprintf("* %s", rich.PlainText))
	}
	return output.String()
}

// https://developers.notion.com/reference/block#numbered-list-item-blocks
type NumberedListItemBlock struct {
	RichText []RichText `json:"rich_text"`
	Color    Color      `json:"color"`
	Children []Block    `json:"children"`
}

func (listItem *NumberedListItemBlock) ToMarkdown() string {
	var output strings.Builder
	for i, rich := range listItem.RichText {
		// TODO fix block numbering
		output.WriteString(fmt.Sprintf("%d. %s", i+1, rich.PlainText))
	}
	return output.String()
}

// https://developers.notion.com/reference/block#to-do-blocks
type TodoBlock struct {
	RichText []RichText `json:"rich_text"`
	Checked  *bool      `json:"checked,omitempty"`
	Color    string     `json:"color"`
	Children []Block    `json:"children"`
}

func (todo *TodoBlock) ToMarkdown() string {
	var output strings.Builder

	var state string
	if *todo.Checked {
		state = "x"
	} else {
		state = " "
	}

	for _, rich := range todo.RichText {
		output.WriteString(fmt.Sprintf("- [%s] %s", state, rich.PlainText))
	}

	return output.String()
}

// https://developers.notion.com/reference/block#code-blocks
type CodeBlock struct {
	RichText []RichText `json:"rich_text"`
	Caption  []RichText `json:"caption"`
	Language string     `json:"language"` // enum
}

func (code *CodeBlock) ToMarkdown() string {
	var output strings.Builder

	output.WriteString(fmt.Sprintf("```%s\n", code.Language))

	for _, rich := range code.RichText {
		output.WriteString(rich.PlainText)
	}

	output.WriteString("```\n")

	return output.String()
}

type RetrieveBlockChildrenResponse struct {
	Object     string   `json:"object"`
	Results    []Block  `json:"results"`
	NextCursor *string  `json:"next_cursor"`
	HasMore    bool     `json:"has_more"`
	Type       string   `json:"type"`
	Block      struct{} `json:"block"`
}
