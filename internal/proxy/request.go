package proxy

import (
	"bufio"
	"net/http"
)

func parseRequest(reader *bufio.Reader) (*http.Request, error) {
	request, err := http.ReadRequest(reader)
	if err != nil {
		return nil, err
	}
	return request, nil
}
