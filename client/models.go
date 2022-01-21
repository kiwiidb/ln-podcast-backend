package client

type PodcastIndexResponse struct {
	Status      string `json:"status"`
	Feeds       []Feed `json:"feeds"`
	Count       int    `json:"count"`
	Query       string `json:"query"`
	Description string `json:"description"`
}

type Feed struct {
	ID                     int    `json:"id"`
	Title                  string `json:"title"`
	URL                    string `json:"url"`
	OriginalURL            string `json:"originalUrl"`
	Link                   string `json:"link"`
	Description            string `json:"description"`
	Author                 string `json:"author"`
	OwnerName              string `json:"ownerName"`
	Image                  string `json:"image"`
	Artwork                string `json:"artwork"`
	LastUpdateTime         int    `json:"lastUpdateTime"`
	LastCrawlTime          int    `json:"lastCrawlTime"`
	LastParseTime          int    `json:"lastParseTime"`
	LastGoodHTTPStatusTime int    `json:"lastGoodHttpStatusTime"`
	LastHTTPStatus         int    `json:"lastHttpStatus"`
	ContentType            string `json:"contentType"`
	ItunesID               int    `json:"itunesId"`
	Generator              string `json:"generator"`
	Language               string `json:"language"`
	Type                   int    `json:"type"`
	Dead                   int    `json:"dead"`
	CrawlErrors            int    `json:"crawlErrors"`
	ParseErrors            int    `json:"parseErrors"`
}
