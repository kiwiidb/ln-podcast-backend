package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func HighlightHandler(w http.ResponseWriter, r *http.Request) {
	payload := []PodcastPayload{
		{
			URL: "https://castleisland.libsyn.com/rss",
			AddressPayload: AddressPayload{
				Callback:    "https://api.flitz.be/lnurl-pay-secondary/FLITZgrDBQdSHU0/kwinten",
				MaxSendable: 400000000,
				MinSendable: 10000,
			},
		},
	}
	err := json.NewEncoder(w).Encode(&payload)
	if err != nil {
		logrus.WithError(err).Error("error encoding payload")
	}
}
