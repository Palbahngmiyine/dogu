package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 정적 파일 서빙
	r.Static("/static", "./web/static")
	r.LoadHTMLGlob("web/template/*")

	// 메인 페이지 라우트
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Dogu",
		})
	})

	r.Run(":8080")
}
