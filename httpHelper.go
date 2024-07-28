package nuget

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func getJson[T any](url string, h *http.Client) (*T, error) {
	r, err := h.Get(url)
	if err != nil {
		return nil, err
	}

	if r.StatusCode != 200 {
		return nil, errors.New("unexpected status code")
	}

	data, _ := io.ReadAll(r.Body)
	r.Body.Close()

	var out T
	if err = json.Unmarshal(data, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
