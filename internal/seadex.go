package internal

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type SeadexEntry struct {
	Title            string `json:"Title"`
	AlternateTitle   string `json:"Alternate Title"`
	BestRelease      string `json:"Best Release"`
	AlternateRelease string `json:"Alternate Release"`
	DualAudio        string `json:"Dual Audio"`
	Notes            string `json:"Notes"`
	Comparisons      string `json:"Comparisons"`
}

//getSeadexJSON A Way to get a kinda better JSON object but have to make your own Live mirror of the sheet using https://github.com/benborgers/opensheet/issues/19
func getSeadexJSON() []SeadexEntry {

	resp, err := http.Get("https://opensheet.elk.sh/12LS5_3snOwqmhWuli1cuAbCGnBdCkvCQY0Xvrwmccis/TV")
	if err != nil {
		log.Panicf("Seadex unavailable")
	}

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Seadex JSON conversion failed")
	}

	// unmarshalling respData to an array of SeadexEntry
	var seadexJson []SeadexEntry
	err = json.Unmarshal(respData, &seadexJson)
	if err != nil {
		log.Fatal("Seadex JSON conversion failed", err)
	}
	return seadexJson
}

// Seadex Running few methods from seadex.go
func Seadex() {
	getSeadexJSON()
}

//getSeadexCSV gets a csv string
//func getSeadexCSV() string {
//	resp, err := http.Get("https://docs.google.com/spreadsheets/d/1emW2Zsb0gEtEHiub_YHpazvBd4lL4saxCwyPhbtxXYM/export?format=csv")
//	if err != nil {
//		log.Fatal("Seadex unavailable")
//	}
//	respData, err := io.ReadAll(resp.Body)
//	if err != nil {
//		log.Fatal("Seadex sheet unreadable")
//	}
//	return string(respData)
//}

// downloadSeadexCSV downloads a seadex doc as a csv
//func downloadSeadexCSV() {
//	resp, err := grab.Get(".", "https://docs.google.com/spreadsheets/d/1emW2Zsb0gEtEHiub_YHpazvBd4lL4saxCwyPhbtxXYM/export?format=csv")
//	if resp == nil || err == nil {
//		log.Fatal("Seadex CSV download failed")
//	}
//}

// getSeadexJSON gets a unusable JSON string
//func getSeadexJSON() string {
//
//	resp, err := http.Get("https://docs.google.com/spreadsheets/d/1emW2Zsb0gEtEHiub_YHpazvBd4lL4saxCwyPhbtxXYM/gviz/tq?tqx=out:json/")
//	if err != nil {
//		log.Panicf("Seadex unavailable")
//	}
//	respData, err := io.ReadAll(resp.Body)
//	if err != nil {
//		log.Fatal("Seadex JSON conversion failed")
//	}
//	return string(respData)
//}
