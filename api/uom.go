package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "pixel-thc-backend-go/db/sqlc"
)

type createUomRequest struct {
	Name         string `json:"name" binding:"required"`
	Abbreviation string `json:"abbreviation" binding:"required"`
}

func (server *Server) createUom(ctx *gin.Context) {
	var req createUomRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUomParams{
		Name:         req.Name,
		Abbreviation: req.Abbreviation,
	}

	uom, err := server.store.CreateUom(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, uom)
}

type getUomRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getUom(ctx *gin.Context) {
	var req getUomRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	uom, err := server.store.GetUom(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, uom)
}

func (server *Server) listUoms(ctx *gin.Context) {
	uoms, err := server.store.ListUoms(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, uoms)
}

type deleteUomRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteUom(ctx *gin.Context) {
	var req deleteUomRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := server.store.DeleteUom(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

// TODO: updateUom
