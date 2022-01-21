package handler

type PodcastPayload struct {
	URL            string         `json:"url"`
	AddressPayload AddressPayload `json:"addressPayload"`
}

type AddressPayload struct {
	Callback    string `json:"callback"`
	MaxSendable int    `json:"maxSendable"`
	MinSendable int    `json:"minSendable"`
}
