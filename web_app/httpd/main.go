package main

import (
  "github.com/gin-gonic/gin"
  "web_app/httpd/handler"
  "web_app/platform/newsfeed"
)

func main() {
  feed := newsfeed.New()
  r := gin.Default()

  r.GET("/ping", handler.PingGet())
  r.GET("/newsfeed", handler.NewsfeedGet(feed))
  r.POST("/newsfeed", handler.NewsfeedPost(feed))

  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}