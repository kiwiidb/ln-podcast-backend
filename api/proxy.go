package handler

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		logrus.Error(err)
	}
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, resp.Body)
}
