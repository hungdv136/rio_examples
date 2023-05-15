package example

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

func CallAPI(ctx context.Context, rootURL string, input map[string]interface{}) (map[string]interface{}, error) {
	bodyBytes, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, rootURL+"/animal", bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	data := map[string]interface{}{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}
