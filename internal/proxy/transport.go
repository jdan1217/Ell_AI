package proxy

import (
	"net/http"
)

func newTransport() *http.Transport {
	t := &http.Transport{}
	return t
}
