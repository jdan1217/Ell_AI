package traffic

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"
)

type TrafficEntry struct {
	Request   RequestData   `json:"request"`
	Response  ResponseData  `json:"response"`
	ID        string        `json:"id"`
	Timestamp time.Time     `json:"timestamp"`
	Duration  time.Duration `json:"duration"`
}

type RequestData struct {
	Method   string              `json:"method"`
	URL      string              `json:"url"`
	Headers  map[string][]string `json:"headers"`
	Body     []byte              `json:"body"`
	Protocol string              `json:"protocol"`
}

type ResponseData struct {
	StatusCode int                 `json:"status"`
	Headers    map[string][]string `json:"headers"`
	Body       []byte              `json:"body"`
	Protocol   string              `json:"protocol"`
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
