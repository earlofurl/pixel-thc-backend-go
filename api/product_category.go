package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type createProductCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) createProductCategory(ctx *gin.Context) {
	var req createProductCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := req.Name
	productCategory, err := server.store.CreateProductCategory(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, productCategory)
}

type getProductCategoryRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getProductCategory(ctx *gin.Context) {
	var req getProductCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	productCategory, err := server.store.GetProductCategory(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, productCategory)
}

func (server *Server) getProductCategories(ctx *gin.Context) {
	productCategories, err := server.store.GetProductCategories(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, productCategories)
}

type deleteProductCategoryRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteProductCategory(ctx *gin.Context) {
	var req deleteProductCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := server.store.DeleteProductCategory(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

// TODO: updateProductCategory
