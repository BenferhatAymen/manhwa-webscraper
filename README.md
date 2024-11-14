**Manhwa Scraper in Go**

### Description
This project is a web scraper written in Go that collects information about Manhwa (Korean comics) from the website [Teamx Manhwa](https://olympustaff.com/). Using the [Colly](https://github.com/gocolly/colly) package, the scraper extracts data on the latest Manhwa titles, chapters, and images, enabling users to easily access and download specific Manhwa content. The project is modular, with functions for retrieving the latest Manhwa, fetching chapters for a specific title, and gathering all images for a particular chapter.

---

### Features
- **Latest Manhwa Titles**: Retrieve the most recent Manhwa titles with their names and links.
- **Chapter List**: For any selected Manhwa, fetch the complete list of available chapters.
- **Chapter Images**: Gather all images from a chosen chapter, enabling easy content download and viewing.
- **Error Handling**: Improved error handling to ensure the scraper continues smoothly even if individual requests fail.

---

### Requirements
- **Go**: Make sure Go is installed on your system.
- **Colly**: The `colly` package for Go is required. You can install it by running:
  ```bash
  go get -u github.com/gocolly/colly/v2
  ```

---

### Installation and Setup
1. Clone this repository:
   ```bash
   git clone https://github.com/BenferhatAymen/manhwa-webscraper.git
   ```
2. Navigate into the project directory:
   ```bash
   cd manhwa-scraper
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```

### Usage
1. **Get Latest Manhwa Titles**
   Retrieve the latest Manhwa titles from the homepage.
   ```go
   manhwas, err := GetLatestManwas()
   if err != nil {
       log.Fatalf("Error fetching manhwas: %v", err)
   }
   for _, manhwa := range manhwas {
       fmt.Println("Title:", manhwa.Title, "Link:", manhwa.Link)
   }
   ```

2. **Get Chapters for a Manhwa**
   For any selected Manhwa, fetch a list of chapters.
   ```go
   chapters, err := GetChaptersFromManhwa(manhwa)
   if err != nil {
       log.Fatalf("Error fetching chapters: %v", err)
   }
   for _, chapter := range chapters {
       fmt.Println("Chapter Title:", chapter.Title, "Link:", chapter.Link)
   }
   ```

3. **Get Chapter Images**
   Retrieve images from a selected chapter.
   ```go
   chapterImages, err := GetChapterImages(chapter)
   if err != nil {
       log.Fatalf("Error fetching images: %v", err)
   }
   for _, img := range chapterImages.Images {
       fmt.Println("Image URL:", img)
   }
   ```

4. **Search for Manhwa by Name**
   Perform a search for a Manhwa title by a Name.
   ```go
   searchResults, err := SearchForMahwa("solo leveling")
   if err != nil {
       log.Fatalf("Error searching for Manhwa: %v", err)
   }
   for _, result := range searchResults {
       fmt.Println("Found Title:", result.Title, "Link:", result.Link)
   }
   ```

---

### Example Workflow
The following example workflow shows how to use the scraper to find a Manhwa, retrieve its chapters, and download the images from a specific chapter.

```go
func main() {
   manhwas, err := GetLatestManwas()
   if err != nil {
       log.Fatalf("Error fetching manhwas: %v", err)
   }

   selectedManhwa := manhwas[0] // Selecting the first Manhwa as an example
   chapters, err := GetChaptersFromManhwa(selectedManhwa)
   if err != nil {
       log.Fatalf("Error fetching chapters: %v", err)
   }

   selectedChapter := chapters[0] // Selecting the first chapter
   chapterImages, err := GetChapterImages(selectedChapter)
   if err != nil {
       log.Fatalf("Error fetching chapter images: %v", err)
   }

   for _, img := range chapterImages.Images {
       fmt.Println("Image URL:", img)
   }
}
```

---

### Error Handling
Each function in this project includes error handling to manage failed requests. The `c.OnError` function in Colly logs any request-level errors, while the return values of functions ensure that all external calls are handled safely.

---

### License
This project is open-source and available under the MIT License. Feel free to modify and use it for personal or commercial purposes.

---

This scraper can be easily extended to download images or save data to a database, making it a powerful tool for those interested in automating the retrieval of online comic content.
