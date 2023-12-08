package main

import (
    "fmt"
    "github.com/gocolly/colly"
    "io"
    "net/http"
    "os"
    "strconv"
    "testing"
    "time"
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
            resp, _ := http.Get(url)
            defer resp.Body.Close()

            file, _ := os.Create(strconv.FormatInt(time.Now().UnixNano(), 10) + ".jpg")
            defer file.Close()

            io.Copy(file, resp.Body)
            fmt.Println("Image saved to", file.Name())
        }
    })

    c.Visit("")
}

func Test02(t *testing.T) {
    start := time.Now()
    ch := make(chan bool)
    for i := 0; i < 10; i++ {
        go parseUrls("https://movie.douban.com/top250?start="+strconv.Itoa(25*i), ch)
    }

    for i := 0; i < 10; i++ {
        <-ch
    }

    elapsed := time.Since(start)
    fmt.Printf("Took %s", elapsed)
}
