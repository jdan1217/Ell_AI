package traffic

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"
)

type TrafficEntry struct {
	Request   RequestData
	Response  ResponseData
	ID        string
	Timestamp time.Time
	Duration  time.Duration
}

type RequestData struct {
	Method   string
	URL      string
	Headers  map[string][]string
	Body     []byte
	Protocol string
}

type ResponseData struct {
	StatusCode int
	Headers    map[string][]string
	Body       []byte
	Protocol   string
}

func NewTrafficEntry(
	req *http.Request,
	resp *http.Response,
	duration time.Duration,
) (TrafficEntry, error) {
	id := make([]byte, 16)
	_, err := rand.Read(id)
	if err != nil {
		return TrafficEntry{}, err
	}
	entry := TrafficEntry{}
	entry.ID = hex.EncodeToString(id)
	entry.Timestamp = time.Now()
	entry.Duration = duration

	requestData, err := NewRequestData(req)
	if err != nil {
		return TrafficEntry{}, err
	}
	entry.Request = requestData
	responseData, err := NewResponseData(resp)
	if err != nil {
		return TrafficEntry{}, err
	}
	entry.Response = responseData
	return entry, nil
}
