package nuget

import (
	"errors"
	"net/http"
)

const serviceIndexUrl = "https://api.nuget.org/v3/index.json"

var ErrServiceNotListed = errors.New("service was not listed by the service index")

type NugetRepository struct {
	httpClient *http.Client
	resources  map[string][]string
}

func NewRepository(h *http.Client) (*NugetRepository, error) {
	index, err := getJson[Services](serviceIndexUrl, h)
	if err != nil {
		return nil, err
	}

	r := make(map[string][]string)
	for _, service := range index.Resources {
		r[service.Type] = append(r[service.Type], service.Id)
	}

	return &NugetRepository{
		httpClient: h,
		resources:  r,
	}, nil
}

func (r *NugetRepository) GetCatalog(cursor Cursor) (*NugetCatalog, error) {
	catalogEndpoint, ok := r.resources["Catalog/3.0.0"]
	if !ok {
		return nil, ErrServiceNotListed
	}

	return &NugetCatalog{
		Cursor:     cursor,
		base:       catalogEndpoint[0],
		httpClient: r.httpClient,
	}, nil
}

func NewCatalog(h *http.Client, cursor Cursor) (*NugetCatalog, error) {
	r, err := NewRepository(h)
	if err != nil {
		return nil, err
	}

	return r.GetCatalog(cursor)
}
