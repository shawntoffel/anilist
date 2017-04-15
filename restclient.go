package anilist

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ResponseError struct {
	Error Error `json:"error"`
}
type Error struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type RestClient interface {
	Get(url string, headers map[string]string, output interface{}) error
	Post(url string, headers map[string]string, body interface{}, output interface{}) error
}

type restClient struct {
	Client *http.Client
}

func NewRestClient() RestClient {
	client := &http.Client{}

	return &restClient{client}
}

func (r *restClient) Get(url string, headers map[string]string, output interface{}) error {

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	req.Header.Add("Content-Type", "application/json")

	response, err := r.Client.Do(req)

	defer response.Body.Close()

	return DecodeJson(response, &output)
}

func (r *restClient) Post(url string, headers map[string]string, body interface{}, output interface{}) error {
	var jsonString = EncodeJson(body)

	reader := bytes.NewReader(jsonString)

	req, err := http.NewRequest("POST", url, reader)

	if err != nil {
		return err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	req.Header.Add("Content-Type", "application/json")

	response, err := r.Client.Do(req)

	defer response.Body.Close()

	return DecodeJson(response, &output)
}

func EncodeJson(body interface{}) []byte {
	bytes, _ := json.Marshal(body)
	return bytes

}

func DecodeJson(r *http.Response, into interface{}) error {

	var decoder = json.NewDecoder(r.Body)

	if r.StatusCode != 200 {
		responseError := ResponseError{}

		err := decoder.Decode(&responseError)

		if err != nil {
			return err

		}

		return errors.New(fmt.Sprintf("AniList API error: %s, %s", responseError.Error.Error, responseError.Error.ErrorDescription))

	}

	return decoder.Decode(&into)

}
