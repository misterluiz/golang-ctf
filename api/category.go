package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/misterluiz/golang-ctf/db/sqlc"
	"github.com/misterluiz/golang-ctf/util"
)

type createCategoryRequest struct {
	UserID      int32  `json:"user_id" binding:"required"`
	Tytle       string `json:"tytle" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (server *Server) createCategory(ctx *gin.Context) {
	errValidateToken := util.GetTokenInHeaderAndVerify(ctx)
	if errValidateToken != nil {
		return
	}

	var req createCategoryRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.CreateCategoryParams{
		UserID:      req.UserID,
		Tytle:       req.Tytle,
		Type:        req.Type,
		Description: req.Description,
	}
	category, err := server.store.CreateCategory(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, category)

}

type getCategoryRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getCategory(ctx *gin.Context) {
	errValidateToken := util.GetTokenInHeaderAndVerify(ctx)
	if errValidateToken != nil {
		return
	}

	var req getCategoryRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	category, err := server.store.GetCategory(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, category)
}

type deleteCategoryRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) deleteCategory(ctx *gin.Context) {
	errValidateToken := util.GetTokenInHeaderAndVerify(ctx)
	if errValidateToken != nil {
		return
	}

	var req deleteCategoryRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	err := server.store.DeleteCategories(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, true)
}

type updateCategoryRequest struct {
	ID          int32  `json:"id" binding:"required"`
	Tytle       string `json:"tytle" `
	Description string `json:"description" `
}

func (server *Server) updateCategory(ctx *gin.Context) {

	errValidateToken := util.GetTokenInHeaderAndVerify(ctx)
	if errValidateToken != nil {
		return
	}

	var req updateCategoryRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.UpdateCategoriesParams{
		ID:          req.ID,
		Tytle:       req.Tytle,
		Description: req.Description,
	}
	category, err := server.store.UpdateCategories(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, category)

}

type getCategoriesRequest struct {
	UserID      int32  `json:"user_id" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Tytle       string `json:"tytle" `
	Description string `json:"description" `
}

func (server *Server) getCategories(ctx *gin.Context) {

	errValidateToken := util.GetTokenInHeaderAndVerify(ctx)
	if errValidateToken != nil {
		return
	}

	var req getCategoriesRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetCategoriesParams{
		UserID:      req.UserID,
		Type:        req.Type,
		Tytle:       req.Tytle,
		Description: req.Description,
	}

	categories, err := server.store.GetCategories(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))

	}

	ctx.JSON(http.StatusOK, categories)
}
