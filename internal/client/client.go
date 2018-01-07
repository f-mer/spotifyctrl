package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/f-mer/spotify-cli/internal/portscanner"
)

type Client struct {
	endpoint url.URL
	csrf     CsrfToken
	oauth    OauthToken
}

type PauseInput struct {
	Pause bool
}

type PlayInput struct {
	Uri     string
	Context string
}

type StatusOutput struct {
	Version       int    `json:"version"`
	ClientVersion string `json:"client_version"`
	Playing       bool   `json:"playing"`
	Shuffle       bool   `json:"shuffle"`
	Repeat        bool   `json:"repeat"`
	PlayEnabled   bool   `json:"play_enabled"`
	PrevEnabled   bool   `json:"prev_enabled"`
	NextEnabled   bool   `json:"next_enabled"`
	Track         struct {
		TrackResource struct {
			Name     string `json:"name"`
			URI      string `json:"uri"`
			Location struct {
				Og string `json:"og"`
			} `json:"location"`
		} `json:"track_resource"`
		ArtistResource struct {
			Name     string `json:"name"`
			URI      string `json:"uri"`
			Location struct {
				Og string `json:"og"`
			} `json:"location"`
		} `json:"artist_resource"`
		AlbumResource struct {
			Name     string `json:"name"`
			URI      string `json:"uri"`
			Location struct {
				Og string `json:"og"`
			} `json:"location"`
		} `json:"album_resource"`
		Length    int    `json:"length"`
		TrackType string `json:"track_type"`
	} `json:"track"`
	Context struct {
	} `json:"context"`
	PlayingPosition float64 `json:"playing_position"`
	ServerTime      int     `json:"server_time"`
	Volume          int     `json:"volume"`
	Online          bool    `json:"online"`
	OpenGraphState  struct {
		PrivateSession bool `json:"private_session"`
	} `json:"open_graph_state"`
	Running bool `json:"running"`
}

func NewClient() (*Client, error) {
	port := portscanner.FirstOpened(4370, 4380)

	oauth, err := getOauthToken()
	if err != nil {
		return nil, err
	}

	csrf, err := getCsrfToken(port)
	if err != nil {
		return nil, err
	}

	endp, err := url.Parse(fmt.Sprintf("http://127.0.0.1:%d", port))
	if err != nil {
		return nil, err
	}

	return &Client{
		endpoint: *endp,
		csrf:     *csrf,
		oauth:    *oauth,
	}, nil
}

func (c *Client) get(path string, qs *url.Values) (*http.Response, error) {
	qs.Set("csrf", c.csrf.Value)
	qs.Set("oauth", c.oauth.Value)

	c.endpoint.RawQuery = qs.Encode()
	c.endpoint.Path = path

	client := &http.Client{}
	req, err := http.NewRequest("GET", c.endpoint.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Origin", "https://open.spotify.com")
	res, err := client.Do(req)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c *Client) Pause(in *PauseInput) (*http.Response, error) {
	res, err := c.get("/remote/pause.json", &url.Values{
		"pause": {strconv.FormatBool(in.Pause)},
	})
	if err != nil {
		return res, err
	}
	defer res.Body.Close()
	return res, nil
}

func (c *Client) Play(in *PlayInput) (*http.Response, error) {
	qs := &url.Values{
		"uri": {in.Uri},
	}
	if len(in.Context) > 0 {
		qs.Set("context", in.Context)
	}
	res, err := c.get("/remote/play.json", qs)
	if err != nil {
		return res, err
	}
	defer res.Body.Close()
	return res, nil
}

func (c *Client) Status() (*StatusOutput, *http.Response, error) {
	res, err := c.get("/remote/status.json", &url.Values{})
	if err != nil {
		return nil, res, err
	}
	defer res.Body.Close()

	out := new(StatusOutput)
	if err := json.NewDecoder(res.Body).Decode(out); err != nil {
		return out, res, err
	}

	return out, res, nil
}
