package anilist

import (
	"fmt"
	"github.com/google/go-querystring/query"
)

var BaseUrl = "https://anilist.co/api/"

type anilist struct {
	Client RestClient
}

type Anilist interface {
	GetAccessToken(AuthenticationRequest) (AuthenticationResponse, error)
	BrowseAnime(accessToken string, request BrowseAnimeRequest) (BrowseAnimeResponse, error)
}

func NewAnilistClient() Anilist {
	restClient := NewRestClient()

	return &anilist{restClient}
}

func (a *anilist) GetAccessToken(request AuthenticationRequest) (AuthenticationResponse, error) {
	response := AuthenticationResponse{}

	var headers map[string]string

	err := a.Client.Post(BaseUrl+"auth/access_token", headers, request, &response)

	return response, err
}

func (a *anilist) BrowseAnime(accessToken string, request BrowseAnimeRequest) (BrowseAnimeResponse, error) {
	response := BrowseAnimeResponse{}

	values, _ := query.Values(request)

	queryString := values.Encode()

	url := BaseUrl + "browse/anime?" + queryString

	fmt.Print(url)

	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + accessToken

	err := a.Client.Get(url, headers, &response)

	return response, err
}
