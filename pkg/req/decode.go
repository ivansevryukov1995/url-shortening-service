package req

import (
	"encoding/json"
	"io"
)

func Decode[T any](body io.ReadCloser) (T, error) {
	var req T
	err := json.NewDecoder(body).Decode(&req)
	if err != nil {
		return req, err
	}

	return req, nil
}
