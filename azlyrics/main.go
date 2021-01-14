package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Response struct {
	artistname string
	featartist string
	songname   string
	lyrics     string
}

func format(text string) string {
	t := strings.ToLower(text)
	t = strings.ReplaceAll(t, " ", "")
	return t
}

func FetchLyrics(artistname, songname string) (response Response) {

	c := colly.NewCollector()

	c.OnHTML(".col-xs-12.col-lg-8.text-center", func(e *colly.HTMLElement) {

		var depth int = 8

		artistname := e.DOM.Find(".lyricsh").Text()
		response.artistname = strings.ReplaceAll(strings.TrimSpace(artistname), "Lyrics", "")

		songname := strings.ReplaceAll(e.DOM.Find(".col-xs-12.col-lg-8.text-center > b").Text(), "\"", "")
		response.songname = strings.TrimSpace(songname)

		feat := e.DOM.Find(".feat")

		if feat.Length() > 0 {
			depth = 10
			response.featartist = feat.Text()
		}

		lyrics := e.DOM.Find(".col-xs-12.col-lg-8.text-center > div:nth-child(" + strconv.Itoa(depth) + ")").Text()
		response.lyrics = strings.TrimSpace(lyrics)

	})

	url := fmt.Sprintf("https://www.azlyrics.com/lyrics/%s/%s.html", artistname, songname)
	c.Visit(url)

	return response
}

func main() {
	artistname := format("50 cent")
	songname := format("candy shop")
	response := FetchLyrics(artistname, songname)
	fmt.Println(response.artistname)
	fmt.Println(response.songname)
}
