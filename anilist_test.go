package anilist

import (
	"fmt"
	"testing"
)

const TestQuery = `query($season: MediaSeason, $seasonYear: Int, $page: Int)  { Page(page: $page, perPage: 10) { pageInfo { total perPage currentPage hasNextPage  } media(season: $season, seasonYear: $seasonYear, format: TV, type: ANIME) { id title { romaji  } siteUrl  }  }  }`

func TestIntegration(t *testing.T) {
	client := New()

	var variables map[string]string
	variables = make(map[string]string)
	variables["season"] = "FALL"
	variables["seasonYear"] = "2018"
	variables["page"] = "1"

	request := Request{
		Query:     TestQuery,
		Variables: variables,
	}

	for {
		response, err := client.Query(request)
		if err != nil {
			t.Error(err.Error())
		}

		t.Log(response)

		if !response.Data.Page.PageInfo.HasNextPage {
			break
		}

		t.Log("Executing next page")

		request.Variables["page"] = fmt.Sprintf("%d", response.Data.Page.PageInfo.CurrentPage+1)
	}
}
