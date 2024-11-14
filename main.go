package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

const BASE_URL = "https://olympustaff.com/"
// manhwa struct
type Manhwa struct {
	Title string
	Link  string
	Image string
}

type Chapter struct {
	Title string
	Link  string
}

type ChapterImages struct {
	Images []string
}



func GetLatestManwas() ([]Manhwa, error) {
	c := colly.NewCollector()
	var manhwas []Manhwa

	c.OnHTML("div.uta", func(e *colly.HTMLElement) {

		image := e.ChildAttr("div.imgu > a > img", "src")
		manhwa := Manhwa{
			Title: e.ChildText("div.info > a > h3"),
			Link:  e.Request.AbsoluteURL(e.ChildAttr("div.info > a", "href")),
			Image: image,
		}
		manhwas = append(manhwas, manhwa)

	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error while visiting %s: %v\n", r.Request.URL, err)
	})

	err := c.Visit(BASE_URL)
	if err != nil {
		return nil, fmt.Errorf("failed to visit %s: %w", BASE_URL, err)
	}

	return manhwas, nil
}

func GetChaptersFromManhwa(manhwa Manhwa) ([]Chapter, error) {
	c := colly.NewCollector()
	var chapters []Chapter
	var pages []string

	c.OnHTML("ul.pagination li > a.page-link", func(e *colly.HTMLElement) {

		pages = append(pages, e.Attr("href"))

	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error while visiting %s: %v\n", r.Request.URL, err)
	})
	err := c.Visit(manhwa.Link)
	if err != nil {
		return nil, fmt.Errorf("failed to visit %s: %w", manhwa.Link, err)

	}
	if len(pages) > 0 {
		pages = pages[:len(pages)-1]
	}
	chapterCollector := colly.NewCollector()

	chapterCollector.OnHTML("div.ts-chl-collapsible-content li", func(e *colly.HTMLElement) {
		chapter := Chapter{
			Title: e.ChildText("div.epl-num"),
			Link:  e.Request.AbsoluteURL(e.ChildAttr("a", "href")),
		}
		chapters = append(chapters, chapter)
	})

	chapterCollector.OnError(func(r *colly.Response, err error) {
		log.Printf("Error while visiting chapter page %s: %v\n", r.Request.URL, err)
	})

	for _, page := range pages {
		if err := chapterCollector.Visit(page); err != nil {
			log.Printf("Error visiting page %s: %v", page, err)
		}
	}

	return chapters, nil
}

func GetChapterImages(chapter Chapter) (ChapterImages, error) {
	c := colly.NewCollector()
	var chapterImages ChapterImages

	c.OnHTML("div.page-break.no-gaps", func(e *colly.HTMLElement) {
		imageSrc := e.ChildAttr("img", "src")
		chapterImages.Images = append(chapterImages.Images, imageSrc)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error while visiting %s: %v\n", r.Request.URL, err)
	})

	err := c.Visit(chapter.Link)
	if err != nil {
		return ChapterImages{}, fmt.Errorf("failed to visit %s: %w", chapter.Link, err)
	}

	return chapterImages, nil
}

func setQuery(query string) string {
	return strings.ReplaceAll(query, " ", "+")
}

func SearchForMahwa(query string) ([]Manhwa, error) {
	c := colly.NewCollector()
	settedQuery := setQuery(query)
	searchUrl := "https://olympustaff.com/?search=" + settedQuery
	var manhwas []Manhwa

	c.OnHTML("div.bsx", func(e *colly.HTMLElement) {
		manhwaLink := e.ChildAttr("a", "href")
		manhwaName := e.ChildAttr("a", "title")
		manhwaImage := e.ChildAttr("a > div.limit > img", "src")
		manhwa := Manhwa{
			Title: manhwaName,
			Link:  manhwaLink,
			Image: manhwaImage,
		}
		manhwas = append(manhwas, manhwa)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error while visiting %s: %v\n", r.Request.URL, err)
	})

	err := c.Visit(searchUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to visit %s: %w", searchUrl, err)
	}

	return manhwas, nil
}

// func main() {
// 	manhwas, err := SearchForMahwa("solo")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(manhwas[2])

// 	chapters , _:= GetChaptersFromManhwa(manhwas[1])

// 	for _, chapter := range chapters {
// 		fmt.Println(chapter)
// 	}

// }
