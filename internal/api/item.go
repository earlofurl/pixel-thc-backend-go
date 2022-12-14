package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/nulls"
	"net/http"
	"pixel-thc-backend-go/internal/db/sqlc"
)

type createItemRequest struct {
	Description string `json:"description" binding:"required"`
	IsUsed      bool   `json:"is_used"`
	ItemTypeId  int64  `json:"item_type_id" binding:"required"`
	StrainId    int64  `json:"strain_id" binding:"required"`
}

func (server *Server) createItem(ctx *gin.Context) {
	var req createItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateItemParams{
		Description: req.Description,
		IsUsed:      req.IsUsed,
		ItemTypeID:  req.ItemTypeId,
		StrainID:    req.StrainId,
	}

	item, err := server.store.CreateItem(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, item)
}

type getItemRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getItem(ctx *gin.Context) {
	var req getItemRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	item, err := server.store.GetItem(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, item)
}

func (server *Server) listItems(ctx *gin.Context) {
	items, err := server.store.ListItems(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, items)
}

type updateItemRequest struct {
	ID          int64        `json:"id" binding:"required"`
	Description nulls.String `json:"description"`
	IsUsed      nulls.Bool   `json:"is_used"`
	ItemTypeId  nulls.Int64  `json:"item_type_id"`
	StrainId    nulls.Int64  `json:"strain_id"`
}

func (server *Server) updateItem(ctx *gin.Context) {
	var req updateItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateItemParams{
		ID:          req.ID,
		Description: req.Description,
		IsUsed:      req.IsUsed,
		ItemTypeID:  req.ItemTypeId,
		StrainID:    req.StrainId,
	}

	item, err := server.store.UpdateItem(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, item)
}

func (server *Server) deleteItem(ctx *gin.Context) {
	var req getItemRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := server.store.DeleteItem(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

// TODO: Add updateItem
