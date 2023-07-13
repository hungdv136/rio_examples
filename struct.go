package example

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type API struct {
	RootURL string
}

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	ID string `json:"id"`
}

func (api *API) Call(ctx context.Context, input Request) (*Response, error) {
	bodyBytes, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, api.RootURL+"/animal", bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		err := errors.New("not found")
		return nil, err
	}

	var data Response
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
