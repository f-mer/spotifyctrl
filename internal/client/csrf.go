package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CsrfToken struct {
	Value string `json:"token"`
}

func getCsrfToken(port int) (*CsrfToken, error) {
	addr := fmt.Sprintf("http://127.0.0.1:%d/simplecsrf/token.json", port)

	client := &http.Client{}
	req, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Origin", "https://open.spotify.com")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	out := new(CsrfToken)
	if err := json.NewDecoder(res.Body).Decode(out); err != nil {
		return out, err
	}

	return out, nil
}
