package http

import (
	"encoding/json"
	"log"
)

type Header map[string][]string

func (h Header) Get(key string) []string {
	return h[key]
}

func (h Header) Set(key string, value []string) {
	h[key] = value
}

func (h Header) Add(key string, value string) {
	h[key] = append(h[key], value)
}

func (h Header) encode() string {
	content, err := json.Marshal(h)
	if err != nil {
		log.Fatal("falied to encode headers ", err)
	}
	return string(content)
}
