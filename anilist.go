package anilist

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const BaseUrl = "https://graphql.anilist.co"

type Anilist interface {
	Query(context.Context, Request) (*Response, error)
}

type anilist struct {
	Client *http.Client
}

func New() Anilist {
	return NewWithClient(&http.Client{})
}

func NewWithClient(client *http.Client) Anilist {
	return &anilist{Client: client}
}

func (a *anilist) Query(ctx context.Context, request Request) (*Response, error) {
	marshalled, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", BaseUrl, bytes.NewBuffer(marshalled))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := a.Client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := &Response{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
