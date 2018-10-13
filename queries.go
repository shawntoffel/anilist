package anilist

const DefaultAnimeForSeasonQuery = `query($season: MediaSeason, $seasonYear: Int, $page: Int)  { Page(page: $page) { pageInfo { total perPage currentPage hasNextPage  } media(season: $season, seasonYear: $seasonYear, format: TV, type: ANIME) { id title { romaji  } siteUrl  }  }  }`
