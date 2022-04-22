package anilist

type Request struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

type Response struct {
	Data   Data    `json:"data"`
	Errors []Error `json:"errors"`
}

type Data struct {
	Page Page `json:"page"`
}

type Media struct {
	Id      int        `json:"id"`
	Title   MediaTitle `json:"title"`
	Type    string     `json:"type"`
	SiteUrl string     `json:"siteUrl"`
	Format  string     `json:"format"`
	Status  string     `json:"status"`
}

type MediaTitle struct {
	Romaji string `json:"romaji"`
}

type Page struct {
	PageInfo PageInfo `json:"pageInfo"`
	Media    []Media
}

type PageInfo struct {
	Total       int  `json:"total"`
	PerPage     int  `json:"perPage"`
	CurrentPage int  `json:"currentPage"`
	HasNextPage bool `json:"hasNextPage"`
}

type Error struct {
	Message   string          `json:"message"`
	Status    int             `json:"status"`
	Locations []ErrorLocation `json:"locations"`
}

type ErrorLocation struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}
