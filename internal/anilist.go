package internal

import (
	"github.com/Dank-del/anilistWrapGO"
	"log"
)

type Poster struct {
	Large, ExtraLarge string
}

type Anime struct {
	poster Poster
	id     int64
	airing []string
}

// getAnilistPoster get posters from AniList
func getAnilistPoster(name string) Poster {
	var animePoster Poster
	anime, err := anilistWrapGO.SearchAnime(name)

	// why is error handling in this language so bad
	if err != nil {
		log.Fatal(err.Error())
	}

	animePoster.Large = anime.Data.Media.CoverImage.Large
	animePoster.ExtraLarge = anime.Data.Media.CoverImage.ExtraLarge

	return animePoster
}

// getAnimeId get anilist id
func getAnimeId(name string) int64 {
	anime, err := anilistWrapGO.SearchAnime(name)
	if err != nil {
		log.Fatal(err.Error())
	}
	animeId := anime.Data.Media.ID
	return animeId
}

// AnilistAnime shitty generator for Anime Struct
func AnilistAnime(name string) Anime {
	var japaneseCartoon Anime

	japaneseCartoon.poster = getAnilistPoster(name)
	japaneseCartoon.id = getAnimeId(name)
	japaneseCartoon.airing = []string{"Destini implement this"}

	return japaneseCartoon
}
