package notion

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"git.sr.ht/~bossley9/sn/pkg/api"
)

type NotionBlocksResponse struct {
	Object     string `json:"object"`
	Results    []NotionBlock
	NextCursor interface{} `json:"next_cursor"`
	HasMore    bool        `json:"has_more"`
	Type       string      `json:"type"`
	Block      struct {
	} `json:"block"`
}

type NotionBlock struct {
	Object string `json:"object"`
	ID     string `json:"id"`
	Parent struct {
		Type   string `json:"type"`
		PageID string `json:"page_id"`
	} `json:"parent"`
	CreatedTime    time.Time `json:"created_time"`
	LastEditedTime time.Time `json:"last_edited_time"`
	CreatedBy      struct {
		Object string `json:"object"`
		ID     string `json:"id"`
	} `json:"created_by"`
	LastEditedBy struct {
		Object string `json:"object"`
		ID     string `json:"id"`
	} `json:"last_edited_by"`
	HasChildren bool                  `json:"has_children"`
	Archived    bool                  `json:"archived"`
	Type        string                `json:"type"`
	Paragraph   *NotionBlockParagraph `json:"paragraph"`
	Todo        *NotionBlockTodo      `json:"to_do"`
	Code        *NotionBlockCode      `json:"code"`
}

type NotionBlockRichText struct {
	Type string `json:"type"`
	Text struct {
		Content string      `json:"content"`
		Link    interface{} `json:"link"`
	} `json:"text"`
	Annotations struct {
		Bold          bool   `json:"bold"`
		Italic        bool   `json:"italic"`
		Strikethrough bool   `json:"strikethrough"`
		Underline     bool   `json:"underline"`
		Code          bool   `json:"code"`
		Color         string `json:"color"`
	} `json:"annotations"`
	PlainText string      `json:"plain_text"`
	Href      interface{} `json:"href"`
}

type NotionBlockParagraph struct {
	RichText []NotionBlockRichText  `json:"rich_text"`
	Color    string                 `json:"color"`
	Children []NotionBlockParagraph `json:"children"`
}

type NotionBlockTodo struct {
	RichText []NotionBlockRichText `json:"rich_text"`
	Checked  bool                  `json:"checked"`
	Color    string                `json:"color"`
	Children []NotionBlockTodo     `json:"children"`
}

type NotionBlockCode struct {
	RichText []NotionBlockRichText `json:"rich_text"`
	Language string                `json:"language"`
}

func (client *Client) GetPageContent(page *NotionPage) []NotionBlock {
	params := map[string]any{}

	block_id := page.ID
	page_size := 100

	url := fmt.Sprintf("https://api.notion.com/v1/blocks/%s/children?page_size=%d", block_id, page_size)

	resp, err := api.Fetch(url, "GET", params, client.Headers)
	if err != nil {
		fmt.Println(err)
		return []NotionBlock{}
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return []NotionBlock{}
	}

	var blocks_response NotionBlocksResponse
	json.Unmarshal(body, &blocks_response)

	return blocks_response.Results
}
