package main

import (
	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/static"
    "net/http"
    "io/ioutil"
    "log"
    "strings"
    "html/template"

    "github.com/PuerkitoBio/goquery"
    "github.com/russross/blackfriday"
)

func initializeRoutes(router *gin.Engine) {

    f, err := ioutil.ReadFile("./_summary")

    if err != nil {
        log.Fatal(err)
    }

    summary_uri := strings.TrimSpace(string(f))

    router.GET("/", func(c *gin.Context) {

        doc, err := goquery.NewDocument(summary_uri)
        if err != nil {
            log.Fatal(err)
        }

        summary_md := doc.Find("div#doc").Text()

        summary_md_arr := strings.Split(summary_md, "---")

        if len(summary_md_arr) >= 3 {
            summary_md = strings.Join(summary_md_arr[2:], "---")
        }

        summary := template.HTML(blackfriday.MarkdownCommon([]byte(summary_md)))

        c.HTML(
            http.StatusOK,
            "index.html",
            gin.H{
                "summary": summary,
                "page": c.Query("uri"),
            },
        )
    })

    // static routes
    router.Use(static.Serve("/static", static.LocalFile("./static", false)))
}
