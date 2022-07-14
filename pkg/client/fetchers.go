package client

import (
	"errors"

	"git.sr.ht/~bossley9/sn/pkg/notion"
)

func fetchPages() ([]notion.Page, error) {
	searchRes, err := notion.Search("", "", 10)
	if err != nil {
		return []notion.Page{}, err
	}

	return searchRes.Results, nil
}

func fetchPageTitle(pageID string) (string, error) {
	propRes, err := notion.RetrievePagePropertyItem(pageID, "title", "", 10)
	if err != nil {
		return "", err
	}

	if len(propRes.Results) == 0 {
		return "", errors.New("no page title found")
	}

	propertyItem := propRes.Results[0]
	if propertyItem.Type != "title" || propertyItem.Title == nil {
		return "", errors.New("invalid page title property returned from the API")
	}

	return propertyItem.Title.PlainText, nil
}

func fetchPageBlocks(pageID string) ([]notion.Block, error) {
	blockRes, err := notion.RetrieveBlockChildren(pageID, "", 100)
	if err != nil {
		return []notion.Block{}, err
	}

	return blockRes.Results, nil
}
