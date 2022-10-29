package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/misterluiz/golang-ctf/db/sqlc"
)

type createAccountRequest struct {
	UserID      int32     `json:"user_id" binding:"required"`
	CategoryID  int32     `json:"category_id" binding:"required"`
	Tytle       string    `json:"tytle" binding:"required"`
	Type        string    `json:"type" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Value       int32     `json:"value" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	var categoryId = req.CategoryID
	var accountType = req.Type

	category, err := server.store.GetCategory(ctx, categoryId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
	}

	var validateCategoryIdDiferentAccountType = category.Type != accountType
	if validateCategoryIdDiferentAccountType {
		ctx.JSON(http.StatusBadRequest, "Account Type is different of category type")
	} else {
		arg := db.CreateAccountParams{
			UserID:      req.UserID,
			CategoryID:  req.CategoryID,
			Tytle:       req.Tytle,
			Type:        req.Type,
			Description: req.Description,
			Value:       req.Value,
			Date:        req.Date,
		}

		account, err := server.store.CreateAccount(ctx, arg)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		ctx.JSON(http.StatusOK, account)
	}

}

type getAccountRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type deleteAccountRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) deleteAccount(ctx *gin.Context) {
	var req deleteAccountRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	err := server.store.DeleteAccount(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, true)
}

type updateAccountRequest struct {
	ID          int32  `json:"id" binding:"required"`
	Tytle       string `json:"tytle"`
	Description string `json:"description"`
	Value       int32  `json:"value" `
}

func (server *Server) updateAccount(ctx *gin.Context) {
	var req updateAccountRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.UpdateAccountParams{
		ID:          req.ID,
		Tytle:       req.Tytle,
		Description: req.Description,
	}
	account, err := server.store.UpdateAccount(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, account)

}

type getAccountsRequest struct {
	UserID      int32     `json:"user_id" binding:"required"`
	Type        string    `json:"type" binding:"required"`
	CategoryID  int32     `json:"category_id"`
	Tytle       string    `json:"tytle"`
	Description string    `json:"description" `
	Date        time.Time `json:"date" `
}

func (server *Server) getAccounts(ctx *gin.Context) {
	var req getAccountsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetAccountsParams{
		UserID: req.UserID,
		Type:   req.Type,
		CategoryID: sql.NullInt32{
			Int32: req.CategoryID,
			Valid: req.CategoryID > 0,
		},
		Tytle:       req.Tytle,
		Description: req.Description,
		Date: sql.NullTime{
			Time:  req.Date,
			Valid: !req.Date.IsZero(),
		},
	}

	accounts, err := server.store.GetAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

type getAccountGraphRequest struct {
	UserID int32  `uri:"user_id" binding:"required"`
	Type   string `uri:"type" binding:"required"`
}

func (server *Server) getAccountGraph(ctx *gin.Context) {
	var req getAccountGraphRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetAccountsGraphParams{
		UserID: req.UserID,
		Type:   req.Type,
	}
	countGraph, err := server.store.GetAccountsGraph(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))

		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, countGraph)
}

type getAccountReports struct {
	UserID int32  `uri:"user_id" binding:"required"`
	Type   string `uri:"type" binding:"required"`
}

func (server *Server) getAccountReports(ctx *gin.Context) {
	var req getAccountGraphRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))

	}

	arg := db.GetAccountsReportsParams{
		UserID: req.UserID,
		Type:   req.Type,
	}
	sumReports, err := server.store.GetAccountsReports(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, sumReports)
}
