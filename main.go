package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MarkdownParser struct {
}

func NewMarkdownParser() *MarkdownParser {
	return &MarkdownParser{}
}

func (p *MarkdownParser) Parse(c *gin.Context) {
	markdown := c.PostForm("markdown")

	// 替换特殊字符
	markdown = escapeSpecialCharacters(markdown)

	// 解析标题
	markdown = parseHeadings(markdown)

	// 解析段落
	markdown = parseParagraphs(markdown)

	htmlResponse := fmt.Sprintf("<html><body>%s</body></html>", markdown)
	c.String(http.StatusOK, htmlResponse)
}

func escapeSpecialCharacters(markdown string) string {
	// 替换特殊字符，如 *、_ 等
	// 例如：将 "*italic*" 替换为 "<em>italic</em>"
	// ...

	return markdown
}

func parseHeadings(markdown string) string {
	// 解析标题语法，如 # Heading
	// 例如：将 "# Heading" 替换为 "<h1>Heading</h1>"
	// ...

	return markdown
}

func parseParagraphs(markdown string) string {
	// 解析段落语法，如普通文本
	// 例如：将 "This is a paragraph" 替换为 "<p>This is a paragraph</p>"
	// ...

	return markdown
}

func main() {
	r := gin.Default()
	parser := NewMarkdownParser()

	r.POST("/parse-markdown", parser.Parse)

	r.Run(":8080")
}
