package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 정적 파일 서빙
	r.Static("/static", "./web/static")

	// HTML 템플릿 로드
	r.LoadHTMLGlob("web/template/*")

	// 라우트 설정
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Dogu",
		})
	})

	r.GET("/api/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "안녕하세요! HTMX가 잘 작동하고 있습니다.")
	})

	r.Run(":8080")
}
