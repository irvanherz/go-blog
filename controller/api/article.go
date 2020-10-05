package controller

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
	ctx.JSON(http.StatusOK, model.Response{
		Status:   http.StatusOK,
		Data:     articles,
		PageInfo: pageInfo,
		Error:    err,
	})
}

// GetID Articles
func (c *ArticleController) GetID(ctx *gin.Context) {
	stringID := ctx.Param("ID")
	ID, err := strconv.ParseInt(stringID, 10, 64)

	article, err := c.ps.ReadByID(ID)

	ctx.JSON(http.StatusOK, model.Response{
		Status: http.StatusOK,
		Data:   article,
		Error:  err,
	})
}

// Post Articles
func (c *ArticleController) Post(ctx *gin.Context) {
	mutationData := model.ArticleMutation{}
	var authorID int64 = 1
	mutationData.AuthorID = &authorID

	if err := ctx.ShouldBind(&mutationData); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Status: http.StatusBadRequest,
			Error:  model.NewRequestError("1234", err.Error()),
		})
		return
	}
	if err := c.ps.Create(&mutationData); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Status: http.StatusBadRequest,
			Error:  err,
		})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Status: http.StatusOK,
		Data:   mutationData,
	})
}

// PutID Articles
func (c *ArticleController) PutID(ctx *gin.Context) {
	stringID := ctx.Param("ID")
	ID, err := strconv.ParseInt(stringID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Status: http.StatusBadRequest,
			Error:  model.NewRequestError("1234", "Invalid ID"),
		})
		return
	}
	mutationData := model.ArticleMutation{}
	if err := ctx.ShouldBind(&mutationData); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Status: http.StatusBadRequest,
			Error:  model.NewRequestError("1234", err.Error()),
		})
		return
	}
	if err := c.ps.UpdateByID(ID, &mutationData); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Status: http.StatusBadRequest,
			Error:  err,
		})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Status: http.StatusOK,
		Data:   mutationData,
	})
}

// DeleteID Articles
func (c *ArticleController) DeleteID(ctx *gin.Context) {
	stringID := ctx.Param("ID")
	ID, _ := strconv.ParseInt(stringID, 10, 64)

	if err := c.ps.DeleteByID(ID); err != nil {
		ctx.JSON(http.StatusOK, model.Response{
			Status: http.StatusOK,
			Error:  err,
		})
	}
	ctx.JSON(http.StatusOK, model.Response{
		Status: http.StatusOK,
		Data:   true,
	})
}
