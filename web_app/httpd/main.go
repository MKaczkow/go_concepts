package main

import (
  "fmt"
  // "github.com/gin-gonic/gin"
  // "web_app/httpd/handler"
  "web_app/platform/newsfeed"
)

func main() {
  // r := gin.Default()

  // r.GET("/ping", handler.PingGet())

  // r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

  feed := newsfeed.New()

  fmt.Println(feed)

  feed.Add(newsfeed.Item{"hello", "How ya doin'?"})

  fmt.Println(feed)

}