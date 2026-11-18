package traffic

import (
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
