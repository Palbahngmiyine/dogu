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

	// HTMX 테스트 API
	r.GET("/api/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "안녕하세요! HTMX가 정상적으로 작동하고 있습니다.")
	})

	r.Run(":8080")
}
