package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func HighlightHandler(w http.ResponseWriter, r *http.Request) {
	payload := []PodcastPayload{
		{
			URL: proxyUrl(r, "https://media.rss.com/flitzpodcast/feed.xml"),
			AddressPayload: AddressPayload{
				Callback:    "https://api.flitz.be/lnurl-pay-secondary/FLITZgrDBQdSHU0/kwinten",
				MaxSendable: 400000000,
				MinSendable: 10000,
			},
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&payload)
	if err != nil {
		logrus.WithError(err).Error("error encoding payload")
	}
}

type PodcastPayload struct {
	URL            string         `json:"url"`
	AddressPayload AddressPayload `json:"addressPayload"`
}

type AddressPayload struct {
	Callback    string `json:"callback"`
	MaxSendable int    `json:"maxSendable"`
	MinSendable int    `json:"minSendable"`
}
