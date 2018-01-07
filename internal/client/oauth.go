package client

import (
	"encoding/json"
	"net/http"
)

type OauthToken struct {
	Value string `json:"t"`
}

func getOauthToken() (*OauthToken, error) {
	res, err := http.Get("https://open.spotify.com/token")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	out := new(OauthToken)
	if err := json.NewDecoder(res.Body).Decode(out); err != nil {
		return out, err
	}

	return out, nil
}
