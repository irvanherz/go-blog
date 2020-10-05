package blog

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/irvanherz/goblog/model"
	"github.com/irvanherz/goblog/service"
)

//ArticleController hehehe
type ArticleController struct {
	ps *service.ArticleService
}

// NewArticleController heheh
func NewArticleController(ps *service.ArticleService) *ArticleController {
	return &ArticleController{ps: ps}
}

// Get Articles
func (c *ArticleController) Get(ctx *gin.Context) {
	var filters model.ArticleQuery
	ctx.ShouldBindQuery(&filters)
	fmt.Printf("%v+", filters)
	articles, pageInfo, err := c.ps.Read(filters)
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title":     "Users",
		"articles":  articles,
		"pageInfo":  pageInfo,
		"error":     err,
		"totalPage": 5,
	})
}

// GetID Article
func (c *ArticleController) GetID(ctx *gin.Context) {
	stringID := ctx.Param("ID")
	ID, err := strconv.ParseInt(stringID, 10, 64)

	article, err := c.ps.ReadByID(ID)

	ctx.HTML(http.StatusOK, "article-detail.html", gin.H{
		"title": "Users",
		"data":  article,
		"error": err,
		"ID":    ID,
	})
}
