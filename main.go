package main

import "github.com/gin-gonic/gin"

func main() {
  r := gin.Default()
  r.GET("/_server/_status", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "name": "hogo",
      "version": "0.0.0",
    })
  })
  r.Run()
}