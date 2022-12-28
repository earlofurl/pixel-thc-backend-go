package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "pixel-thc-backend-go/db/sqlc"
)

type createStrainRequest struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
}

func (server *Server) createStrain(ctx *gin.Context) {
	var req createStrainRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateStrainParams{
		Name: req.Name,
		Type: req.Type,
	}

	strain, err := server.store.CreateStrain(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, strain)
}

type getStrainRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getStrain(ctx *gin.Context) {
	var req getStrainRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	strain, err := server.store.GetStrain(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, strain)
}

func (server *Server) listStrains(ctx *gin.Context) {
	strains, err := server.store.ListStrains(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, strains)
}

type deleteStrainRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteStrain(ctx *gin.Context) {
	var req deleteStrainRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := server.store.DeleteStrain(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
