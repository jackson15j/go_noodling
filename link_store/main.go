package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Page struct {
	Title string
	Lines []Line
}

type Line struct {
	Url         string
	DisplayText string
	Date        string
}

func loadPageByLine() (*Page, error) {
	filename := "data/urls.txt"
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	var lines []Line
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("read file line error: %v", err)
			break
		}
		fields := strings.Split(line, ",")
		if len(fields) != 3 {
			lines = append(lines, Line{fields[0], fields[0], "??"})
		} else {
			lines = append(lines, Line{fields[0], fields[1], fields[2]})
		}
	}

	// Reverse the existing file order: newest line first.
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}

	// log.Print(lines)
	title := "Link Store"
	return &Page{Title: title, Lines: lines}, nil
}

func urlHandler(w http.ResponseWriter, r *http.Request) {
	p, _ := loadPageByLine()
	templates := template.Must(template.ParseFiles("index.html"))
	err := templates.ExecuteTemplate(w, "index.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// TODO: Write inbound URLS to top of: `urls.txt`, as: `url\n` or `url,displayText,date\n`

type StoreUrl struct {
	Url         string
	DisplayText string
}

// API to POST URLS to.
// Looking into the standard library for a change:
// https://www.jetbrains.com/guide/go/tutorials/rest_api_series/stdlib/
// https://gobyexample.com/http-server
func storeUrlHandler(w http.ResponseWriter, r *http.Request) {
	var storeUrl StoreUrl
	if err := json.NewDecoder(r.Body).Decode(&storeUrl); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	filename := "data/urls.txt"
	f, err := os.OpenFile(
		filename,
		os.O_APPEND|os.O_WRONLY|os.O_CREATE,
		0644,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer f.Close()

	// NOTE: Apple Shortcuts injects newline into the values on POST.
	url := strings.TrimSuffix(storeUrl.Url, "\n")
	displayText := strings.TrimSuffix(storeUrl.DisplayText, "\n")
	date := time.Now().UTC().Format(time.DateTime)
	line := fmt.Sprintf("%s,%s,%s\n", url, displayText, date)
	log.Print(line)
	f.WriteString(line)

	w.WriteHeader(http.StatusOK)
}

func main() {
	log.Print("main")
	http.HandleFunc("/", urlHandler)
	http.HandleFunc("/api/store-url", storeUrlHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
