package main

import (
    "fmt"
    "net/http"
    "github.com/PuerkitoBio/goquery"
)

func main() {
    resp, err := http.Get("https://example.com")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        panic(err)
    }

    doc.Find("h1").Each(func(i int, s *goquery.Selection) {
        fmt.Println(s.Text())
    })
}