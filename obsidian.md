# HYRI
## internal
### anilist.go
- Poster *struct*
	- Large *string*
	-  ExtraLarge *string*
- Anime *struct*
	- poster *Poster* 
	- id *int64*
	- airing *[]string*
- getAnilistPoster() -> Poster
- getAnimeId() -> int64
- AnilistAnime() -> Anime
### db.go
### seadex.go
- SeadexEntry *struct*
	- Title string `json:"Title"`
	- AlternateTitle string `json:"Alternate Title"`
	- BestRelease string `json:"Best Release"`
	- AlternateRelease string `json:"Alternate Release"`
	- DualAudio string `json:"Dual Audio"`
	- Notes string `json:"Notes"`
	- Comparisons string `json:"Comparisons"`
- Seadex()
- getSeadexJSON() -> []SeadexEntry
## main.go
