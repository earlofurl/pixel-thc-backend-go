package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "pixel-thc-backend-go/db/sqlc"
)

type createItemTypeRequest struct {
	ProductForm       string `json:"product_form" binding:"required"`
	ProductModifier   string `json:"product_modifier" binding:"required"`
	UomDefault        int64  `json:"uom_default" binding:"required"`
	ProductCategoryId int64  `json:"product_category_id" binding:"required"`
}

func (server *Server) createItemType(ctx *gin.Context) {
	var req createItemTypeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateItemTypeParams{
		ProductForm:       req.ProductForm,
		ProductModifier:   req.ProductModifier,
		UomDefault:        req.UomDefault,
		ProductCategoryID: req.ProductCategoryId,
	}

	itemType, err := server.store.CreateItemType(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, itemType)
}

type getItemTypeRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getItemType(ctx *gin.Context) {
	var req getItemTypeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	itemType, err := server.store.GetItemType(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, itemType)
}

func (server *Server) listItemTypes(ctx *gin.Context) {
	itemTypes, err := server.store.ListItemTypes(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, itemTypes)
}

func (server *Server) deleteItemType(ctx *gin.Context) {
	var req getItemTypeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := server.store.DeleteItemType(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

// TODO: UpdateItemType
