package main

import (
	"github.com/PuerkitoBio/goquery"
	"html"
	"regexp"
	"strings"
)

const (
	lyricsScriptSelector = "body > div > div.center > script:nth-child(2)"
	titleSelector        = "body > div > div.center > div > div > div.song_info.clearfix > div.person_list_and_other > div > h1"
)

var (
	lyricsRegexp = regexp.MustCompile("(?m)var lyrics = '(.+?)'")
)

type Song struct {
	Title, Lyrics string
}

func extractTitle(doc *goquery.Document) (string, error) {
	titleHtml, err := doc.Find(titleSelector).Html()
	if err != nil {
		return "", err
	}
	return html.UnescapeString(titleHtml), nil
}

func extractLyrics(doc *goquery.Document) (string, error) {
	scriptHtml, err := doc.Find(lyricsScriptSelector).Html()
	if err != nil {
		return "", err
	}
	script := html.UnescapeString(scriptHtml)
	match := lyricsRegexp.FindStringSubmatch(script)
	if len(match) == 0 {
		return "", nil
	}
	lyrics := html.UnescapeString(match[1])
	return strings.Replace(lyrics, "<br>", "\n", -1), nil
}

func fetchSong(url string) (*Song, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	title, err := extractTitle(doc)
	if err != nil {
		return nil, err
	}
	lyrics, err := extractLyrics(doc)
	if err != nil {
		return nil, err
	}

	return &Song{
		Title:  title,
		Lyrics: lyrics,
	}, nil
}
