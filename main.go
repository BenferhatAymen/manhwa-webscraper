package scraper

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

const BASE_URL = "https://olympustaff.com/"

type Manhwa struct {
	Title string
	Link  string
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

	c.OnHTML("div.info > a", func(e *colly.HTMLElement) {

		manhwa := Manhwa{
			Title: e.ChildText("h3"),
			Link:  e.Request.AbsoluteURL(e.Attr("href")),
		}

		manhwas = append(manhwas, manhwa)
	})
	c.Visit(BASE_URL)

	return manhwas, nil
}

func GetChaptersFromManhwa(manhwa Manhwa) ([]Chapter, error) {
	c := colly.NewCollector()

	var chapters []Chapter

	c.OnHTML("div.ts-chl-collapsible-content li", func(e *colly.HTMLElement) {
		chapter := Chapter{
			Title: e.ChildText("div.epl-num"),
			Link:  e.Request.AbsoluteURL(e.ChildAttr("a", "href")),
		}

		chapters = append(chapters, chapter)
	})
	c.Visit(manhwa.Link)

	return chapters, nil
}

func GetChapterImages(chapter Chapter) (ChapterImages, error) {
	c := colly.NewCollector()

	var chapterImages ChapterImages

	c.OnHTML("div.page-break.no-gaps", func(e *colly.HTMLElement) {
		imageSrc := e.ChildAttr("img", "src")
		chapterImages.Images = append(chapterImages.Images, imageSrc)
	})

	c.Visit(chapter.Link)

	return chapterImages, nil
}

func setQuery(query string) string {
	query = strings.ReplaceAll(query, " ", "+")
	return query
}

func SearchForMahwa(query string) ([]Manhwa, error) {
	c := colly.NewCollector()

	settedQuery := setQuery(query)

	searchUrl := "https://olympustaff.com/?search=" + settedQuery
	var manhwas []Manhwa

	c.OnHTML("div.bsx", func(e *colly.HTMLElement) {

		manhwaLink := e.ChildAttr("a", "href")
		manhwaName := e.ChildAttr("a", "title")

		manhwa := Manhwa{
			Title: manhwaName,
			Link:  manhwaLink,
		}
		manhwas = append(manhwas, manhwa)

	})

	c.Visit(searchUrl)
	return manhwas, nil

}
