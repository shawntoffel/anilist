package anilist

import (
	"context"
	"fmt"
	"testing"
)

func TestIntegration(t *testing.T) {
	client := New()

	variables := map[string]interface{}{}
	variables["season"] = "FALL"
	variables["seasonYear"] = "2018"
	variables["page"] = "1"
	variables["formats"] = []string{"TV", "ONA"}

	request := Request{
		Query:     DefaultAnimeForSeasonQuery,
		Variables: variables,
	}

	for {
		response, err := client.Query(context.Background(), request)
		if err != nil {
			t.Error(err.Error())
		}

		t.Logf("%+v", response)

		if !response.Data.Page.PageInfo.HasNextPage {
			break
		}

		t.Log("Executing next page")

		request.Variables["page"] = fmt.Sprintf("%d", response.Data.Page.PageInfo.CurrentPage+1)
	}
}
