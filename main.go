package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

const BASE_URL = "https://olympustaff.com/"

type Manhwa struct {
	Title string
	Link  string
}

type Chapter struct {
	Title  string
	Link   string
	Images []string
}

func main() {
	c := colly.NewCollector()

	firstLinkVisited := false
	firstChapterVisited := false

	var chapter Chapter

	c.OnHTML("div.info > a", func(e *colly.HTMLElement) {
		if firstLinkVisited {
			return
		}
		firstLinkVisited = true

		manhwa := Manhwa{
			Title: e.ChildText("h3"),
			Link:  e.Request.AbsoluteURL(e.Attr("href")),
		}
		fmt.Printf("Manhwa Title: %s\nManhwa Link: %s\n\n", manhwa.Title, manhwa.Link)

		e.Request.Visit(manhwa.Link)
	})

	c.OnHTML("div.ts-chl-collapsible-content li", func(e *colly.HTMLElement) {
		if firstChapterVisited {
			return
		}
		firstChapterVisited = true

		chapter = Chapter{
			Title: e.ChildText("div.epl-num"),
			Link:  e.Request.AbsoluteURL(e.ChildAttr("a", "href")),
		}
		fmt.Printf("Chapter Title: %s\nChapter Link: %s\n\n", chapter.Title, chapter.Link)

		e.Request.Visit(chapter.Link)
	})

	c.OnHTML("div.page-break.no-gaps", func(e *colly.HTMLElement) {
		imageSrc := e.ChildAttr("img", "src")
		chapter.Images = append(chapter.Images, imageSrc)
		fmt.Printf("Added Image: %s\n", imageSrc)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL.String())
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Printf("\nFinal Chapter Data: %+v\n", chapter)
	})

	c.Visit(BASE_URL)
}
