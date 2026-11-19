package proxy

import (
	"net/http"
)

func newTransport() *http.Transport {
	t := &http.Transport{}
	return t
}

func forwardRequest(req *http.Request, transport *http.Transport) (*http.Response, error) {
	return transport.RoundTrip(req)
}
