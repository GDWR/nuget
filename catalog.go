package nuget

import (
	"net/http"
	"time"
)

type Cursor = time.Time

type NugetCatalog struct {
	Cursor Cursor

	base       string
	httpClient *http.Client
}

func (c *NugetCatalog) Catalog() (*Catalog, error) {
	return getJson[Catalog](c.base, c.httpClient)
}

func (c *NugetCatalog) Pages() ([]CatalogItem, error) {
	catalog, err := c.Catalog()
	if err != nil {
		return nil, err
	}

	pages := make([]CatalogItem, 0)
	for _, item := range catalog.Items {
		timestamp, err := time.Parse(TimestampFormat, item.CommitTimestamp)
		if err != nil {
			return nil, err
		}

		if timestamp.After(c.Cursor) {
			pages = append(pages, item)
		}
	}

	return pages, nil
}

func (c *NugetCatalog) Leaves() ([]CatalogItem, error) {
	pages, err := c.Pages()
	if err != nil {
		return nil, err
	}

	leaves := make([]CatalogItem, 0)
	for _, page := range pages {
		page, err := getJson[Catalog](page.Id, c.httpClient)
		if err != nil {
			return nil, err
		}

		for _, item := range page.Items {
			timestamp, err := time.Parse(TimestampFormat, item.CommitTimestamp)
			if err != nil {
				return nil, err
			}

			if timestamp.After(c.Cursor) {
				leaves = append(leaves, item)
			}
		}
	}

	return leaves, nil
}

func (c *NugetCatalog) StreamLeaves() <-chan CatalogLeaf {
	ch := make(chan CatalogLeaf)

	go func() {
		defer close(ch)

		pages, err := c.Pages()
		if err != nil {
			return
		}

		for _, page := range pages {
			page, err := getJson[Catalog](page.Id, c.httpClient)
			if err != nil {
				return
			}

			for _, item := range page.Items {
				timestamp, err := time.Parse(TimestampFormat, item.CommitTimestamp)
				if err != nil {
					return
				}

				if timestamp.After(c.Cursor) {
					leaf, err := getJson[CatalogLeaf](item.Id, c.httpClient)
					if err != nil {
						return
					}

					ch <- *leaf
				}
			}
		}
	}()

	return ch
}
