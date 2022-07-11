package notion

import (
	"fmt"
	"strings"
)

func (client *Client) PageToText(page *NotionPage) string {
	var output strings.Builder

	if page == nil {
		return output.String()
	}

	blocks := client.GetPageContent(page)

	for _, block := range blocks {
		output.WriteString(fmt.Sprintf("%s\n", BlockToText(&block)))
	}

	return output.String()
}

func BlockToText(block *NotionBlock) string {
	switch block.Type {
	case "paragraph":
		return BlockParagraphToText(block.Paragraph)
	case "to_do":
		return BlockTodoToText(block.Todo)
	case "code":
		return BlockCodeToText(block.Code)
	default:
		return fmt.Sprintf("%s UNIMPLEMENTED\n", block.Type)
	}
}

func BlockParagraphToText(paragraph *NotionBlockParagraph) string {
	var output strings.Builder
	for _, rich := range paragraph.RichText {
		output.WriteString(rich.Text.Content)
	}
	return output.String()
}

func BlockTodoToText(todo *NotionBlockTodo) string {
	var output strings.Builder

	if todo.Checked {
		output.WriteString("- [x] ")
	} else {
		output.WriteString("- [ ] ")
	}

	for _, rich := range todo.RichText {
		output.WriteString(rich.Text.Content)
	}

	return output.String()
}

func BlockCodeToText(code *NotionBlockCode) string {
	var output strings.Builder

	output.WriteString(fmt.Sprintf("```%s\n", code.Language))
	for _, rich := range code.RichText {
		output.WriteString(rich.Text.Content)
	}
	output.WriteString("```")

	return output.String()
}
