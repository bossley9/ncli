package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"git.sr.ht/~bossley9/sn/pkg/api"
)

const NUM_RESULTS = 50

type NotionSearchResponse struct {
	Object         string `json:"object"`
	Results        []NotionPage
	NextCursor     interface{} `json:"next_cursor"`
	HasMore        bool        `json:"has_more"`
	Type           string      `json:"type"`
	PageOrDatabase struct {
	} `json:"page_or_database"`
}

type NotionPage struct {
	Object         string    `json:"object"`
	ID             string    `json:"id"`
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
	Cover struct {
		Type     string `json:"type"`
		External struct {
			URL string `json:"url"`
		} `json:"external"`
	} `json:"cover"`
	Icon struct {
		Type  string `json:"type"`
		Emoji string `json:"emoji"`
	} `json:"icon"`
	Parent struct {
		Type       string `json:"type"`
		DatabaseID string `json:"database_id"`
	} `json:"parent"`
	Archived   bool `json:"archived"`
	Properties struct {
		StoreAvailability struct {
			ID string `json:"id"`
		} `json:"Store availability"`
		FoodGroup struct {
			ID string `json:"id"`
		} `json:"Food group"`
		Price struct {
			ID string `json:"id"`
		} `json:"Price"`
		ResponsiblePerson struct {
			ID string `json:"id"`
		} `json:"Responsible Person"`
		LastOrdered struct {
			ID string `json:"id"`
		} `json:"Last ordered"`
		CostOfNextTrip struct {
			ID string `json:"id"`
		} `json:"Cost of next trip"`
		Recipes struct {
			ID string `json:"id"`
		} `json:"Recipes"`
		Description struct {
			ID string `json:"id"`
		} `json:"Description"`
		InStock struct {
			ID string `json:"id"`
		} `json:"In stock"`
		NumberOfMeals struct {
			ID string `json:"id"`
		} `json:"Number of meals"`
		Photo struct {
			ID string `json:"id"`
		} `json:"Photo"`
		Name struct {
			ID string `json:"id"`
		} `json:"Name"`
	} `json:"properties"`
	URL string `json:"url"`
}

func (client *Client) search(query string) (*NotionSearchResponse, error) {
	params := map[string]any{
		"query": query,
		"sort": map[string]string{
			"direction": "ascending",
			"timestamp": "last_edited_time",
		},
		"filter": map[string]string{
			"property": "object",
			"value":    "page",
		},
		"page_size": NUM_RESULTS,
	}

	resp, err := api.Fetch("https://api.notion.com/v1/search", "POST", params, client.Headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var search_response NotionSearchResponse
	json.Unmarshal(body, &search_response)

	return &search_response, nil
}

func (client *Client) GetPages() []NotionPage {
	searchResponse, err := client.search("")
	if err != nil {
		fmt.Println(err)
		return []NotionPage{}
	}
	return searchResponse.Results
}
