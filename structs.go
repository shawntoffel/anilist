package anilist

type AuthenticationRequest struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type AuthenticationResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Expires     int    `json:"expires"`
}

type BrowseAnimeRequest struct {
	Year          string   `url:"year,omitempty"`
	Season        string   `url:"season,omitempty"`
	Type          string   `url:"type,omitempty"`
	Status        string   `url:"status,omitempty"`
	Genres        []string `url:"genres,omitempty"`
	GenresExclude []string `url:"genres_exclude,omitempty"`
	Sort          string   `url:"sort,omitempty"`
	AiringData    bool     `url:"airing_data,omitempty"`
	FullPage      bool     `url:"full_page,omitempty"`
	Page          int      `url:"page,omitempty"`
}

type BrowseAnimeResponse []Anime

type Anime struct {
	Id            int      `json:"id"`
	TitleRomaji   string   `json:"title_romaji"`
	Type          string   `json:"type"`
	ImageUrlMed   string   `json:"image_url_med"`
	ImageUrlSml   string   `json:"image_url_sml"`
	TitleJapanese string   `json:"title_japanese"`
	TitleEnglish  string   `json:"title_english"`
	Synonyms      []string `json:"synonyms"`
	ImageUrlLge   string   `json:"image_url_lge"`
	AiringStatus  string   `json:"airing_status"`
	AverageScore  float32  `json:"average_score"`
	TotalEpisodes int      `json:"total_episodes"`
	Adult         bool     `json:"adult"`
	Popularity    int      `json:"popularity"`
}
