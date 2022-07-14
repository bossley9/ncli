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

func fetchPageTitle(respCh chan<- Promise[string], pageID string) {
	propRes, err := notion.RetrievePagePropertyItem(pageID, "title", "", 10)
	if err != nil {
		respCh <- Promise[string]{"", nil}
	}

	if len(propRes.Results) == 0 {
		err := errors.New("no page title found")
		respCh <- Promise[string]{"", err}
	}

	propertyItem := propRes.Results[0]
	if propertyItem.Type != "title" || propertyItem.Title == nil {
		err := errors.New("invalid page title property returned from the API")
		respCh <- Promise[string]{"", err}
	}

	respCh <- Promise[string]{propertyItem.Title.PlainText, nil}
}

func fetchPageBlocks(respCh chan<- Promise[[]notion.Block], pageID string) {
	res, err := notion.RetrieveBlockChildren(pageID, "", 100)
	if err != nil {
		respCh <- Promise[[]notion.Block]{[]notion.Block{}, err}
	}

	respCh <- Promise[[]notion.Block]{res.Results, err}
}
