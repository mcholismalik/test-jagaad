package util

import (
	"encoding/json"
	"net/http"
)

func FetchJSON(url string, data interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(&data)
}
