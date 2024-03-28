package http

import (
	"bytes"
	"errors"
	"fmt"
)

type HTTPRequest struct {
	RequestLine []byte
	Headers     map[string]string
	Body        []byte
}

func (s HTTPRequest) String() string {
	return fmt.Sprintf("RequestLine: %s\nHeaders: %v\nBody: %s\n", s.RequestLine, s.Headers, s.Body)
}

func NewHTTPRequest(data []byte) (*HTTPRequest, error) {
	req := &HTTPRequest{
		Headers: make(map[string]string),
	}
	req.decode(data)
	return req, nil
}

func (r *HTTPRequest) decode(data []byte) error {
	// Assume data is a valid HTTP request
	requestLine, rest, found := bytes.Cut(data, []byte("\r\n"))
	if !found {
		return errors.New("invalid request, missing request line")
	}
	r.RequestLine = requestLine

	// Parse headers
	headers, body, found := bytes.Cut(rest, []byte("\r\n\r\n"))
	if !found {
		return errors.New("invalid request, missing headers")
	}
	for _, header := range bytes.Split(headers, []byte("\r\n")) {
		key, value, found := bytes.Cut(header, []byte(":"))
		if !found {
			return errors.New("invalid header format")
		}
		r.Headers[string(key)] = string(value)
	}

	// Parse Body
	r.Body = body
	return nil

}
