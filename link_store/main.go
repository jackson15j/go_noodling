package main

import (
	"bufio"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Lines []string
}

func loadPageByLine() (*Page, error) {
	filename := "urls.txt"
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	r := bufio.NewReader(f)
	var lines []string
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("read file line error: %v", err)
			break
		}
		lines = append(lines, line)
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

func main() {
	log.Print("main")
	http.HandleFunc("/", urlHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
