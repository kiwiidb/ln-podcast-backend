package client

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

const (
	searchEndpoint = "/api/1.0/search/byterm"
)

//PodCastIndexClient makes requests to podcast index API
type PodCastIndexClient struct {
	ApiKey    string
	UserAgent string
	ApiSecret string
	Host      string
}

func (p *PodCastIndexClient) SetAuthHeaders(req *http.Request) {
	// prep for crypto
	now := time.Now()
	var apiHeaderTime string = strconv.FormatInt(now.Unix(), 10)
	var data4Hash string = p.ApiKey + p.ApiSecret + apiHeaderTime
	// ======== Hash them to get the Authorization token ========
	h := sha1.New()
	h.Write([]byte(data4Hash))
	hash := h.Sum(nil)
	var hashString string = fmt.Sprintf("%x", hash)
	req.Header.Set("User-Agent", "SuperPodcastPlayer/1.8")
	req.Header.Set("X-Auth-Date", apiHeaderTime)
	req.Header.Set("X-Auth-Key", p.ApiKey)
	req.Header.Set("Authorization", hashString)
}

//PodcastPayloadExternal contains the url for the RSS feed
//and the LN address payload
type PodcastPayloadExternal struct {
	URL            string `json:"url"`
	AddressPayload struct {
		Callback    string `json:"callback"`
		MaxSendable int    `json:"maxSendable"`
		MinSendable int    `json:"minSendable"`
	} `json:"addressPayload"`
}

func NewPodClient() *PodCastIndexClient {
	apiKey := os.Getenv("PI_API_KEY")
	apiSecret := os.Getenv("PI_API_SECRET")

	return &PodCastIndexClient{
		ApiKey:    apiKey,
		UserAgent: "test",
		ApiSecret: apiSecret,
		Host:      "https://api.podcastindex.org/",
	}
}

func (p *PodCastIndexClient) Search(query string) (result *PodcastIndexResponse, err error) {
	// ======== Send the request and collect/show the results ========
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", p.Host, searchEndpoint), nil)
	if err != nil {
		return nil, err
	}
	v := url.Values{}
	v.Add("q", query)
	req.URL.RawQuery = v.Encode()
	p.SetAuthHeaders(req)
	res, getErr := http.DefaultClient.Do(req)
	if getErr != nil {
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}
	result = &PodcastIndexResponse{}
	err = json.NewDecoder(res.Body).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
