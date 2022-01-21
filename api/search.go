package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/flitz-be/ln-podcast-backend/client"
	"github.com/mcnijman/go-emailaddress"
	"github.com/sirupsen/logrus"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	feeds, err := client.NewPodClient().Search(query)
	if err != nil {
		logrus.WithError(err).Errorf("Error searching with query %s", query)
	}
	payload := []PodcastPayload{}
	for _, f := range feeds.Feeds {
		payload = append(payload, PodcastPayload{
			URL:            proxyUrl(r, f.URL),
			AddressPayload: parseForLNAddress(f.Description),
		})
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&payload)
	if err != nil {
		logrus.WithError(err).Error("error encoding payload")
	}
}

func proxyUrl(r *http.Request, in string) (out string) {
	return fmt.Sprintf("https://%s/%s?url=%s", r.Host, strings.Replace(r.URL.Path, "search", "proxy", 1), url.QueryEscape(in))
}
func parseForLNAddress(input string) (result AddressPayload) {
	emails := emailaddress.Find([]byte(input), false)

	for _, e := range emails {
		url := constructLNURL(e.LocalPart, e.Domain)
		resp, err := http.Get(url)
		if err != nil {
			return result
		}
		payload := &LNURLPayResponse{}
		err = json.NewDecoder(resp.Body).Decode(&payload)
		if err != nil {
			return result
		}
		if payload.Callback != "" {
			return AddressPayload{
				Callback:    payload.Callback,
				MaxSendable: payload.MaxSendable,
				MinSendable: payload.MinSendable,
			}
		}
	}

	return result
}

func constructLNURL(user, host string) (result string) {
	return fmt.Sprintf("https://%s/.well-known/lnurlp/%s", host, user)
}

type LNURLPayResponse struct {
	Callback    string `json:"callback"`
	MaxSendable int    `json:"maxSendable"`
	Metadata    string `json:"metadata"`
	MinSendable int    `json:"minSendable"`
	Tag         string `json:"tag"`
}
