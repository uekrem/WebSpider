package main

import (
	"database/sql"
	"log"
	"net/http"
	"strings"
	"fmt"
    "github.com/PuerkitoBio/goquery"
    _ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "webSpider.db")
	if err != nil {
		log.Fatal(err)
	}

	insertStatement := "DROP TABLE IF EXISTS sites"
	db.Exec(insertStatement)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS sites (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		url TEXT,
		title TEXT,
		content TEXT
	)`)
	if err != nil {
		log.Fatal(err)
	}

	var startURL string
	fmt.Print("Provide an address with HTTP or HTTPS protocol:")
	fmt.Scanln(&startURL)
	spider(startURL, db, startURL)
	
	defer db.Close()
}

func spider(url string, db *sql.DB, hide string) {
	

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching URL %s: %v", url, err)
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("Error parsing HTML from URL %s: %v", url, err)
		return
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		title := s.Text()
		content := doc.Find("body").Text()
		cleanedContent := strings.Map(func(r rune) rune {
			if r >= 32 || r == '\n' || r == '\t' || r == '\r' {
				return r
			}
			return -1
		}, content)

		var count int
		query := "SELECT COUNT(*) FROM sites WHERE url = ?"
		db.QueryRow(query, href).Scan(&count)

		if strings.HasPrefix(href, hide) && count == 0{
			insertStatement := "INSERT INTO sites (url, title, content) VALUES (?, ?, ?)"
			db.Exec(insertStatement, href, title, cleanedContent)
			//log.Printf("%v	%v	%v\n\n\n",href ,title , cleanedContent)
			//veriler database yüklenirken incelenebilir
			spider(href, db, hide)
		}
	})
	log.Printf("webSpider a transaction completed")
	defer resp.Body.Close()
}

