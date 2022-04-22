package anilist

const DefaultAnimeForSeasonQuery = `query($season: MediaSeason, $seasonYear: Int, $formats: [MediaFormat], $page: Int)  { Page(page: $page) { pageInfo { total perPage currentPage hasNextPage  } media(season: $season, seasonYear: $seasonYear, format_in: $formats, type: ANIME) { id title { romaji  } siteUrl format status }  }  }`
