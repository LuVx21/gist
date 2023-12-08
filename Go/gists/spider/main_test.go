package main

import (
    "fmt"
    "github.com/gocolly/colly"
    "testing"
)

func Test01(t *testing.T) {
    c := colly.NewCollector()

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL.String())
    })

    c.OnResponse(func(r *colly.Response) {
        fmt.Println("Visited", r.Request.URL.String())
    })

    c.OnError(func(r *colly.Response, err error) {
        fmt.Println("Error:", err)
    })

    c.OnHTML("img", func(e *colly.HTMLElement) {
        url := e.Attr("src")
        if url != "" {
            fmt.Println("Found image:", url)
        }
    })

    c.Visit("")
}
