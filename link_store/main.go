package main

import (
	"bufio"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
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
	filename := "urls.txt"
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
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

// TODO: API to POST URLS to.

func main() {
	log.Print("main")
	http.HandleFunc("/", urlHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
