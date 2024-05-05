package card

import (
	"encoding/json"
	"io"
	"net/http"
)

type Response struct {
	Data []Card `json:"data"`
}

func FetchYuGiOhCards(url string) ([]Card, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
