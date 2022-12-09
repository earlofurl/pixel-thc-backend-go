package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/nulls"
	"net/http"
	db "pixel-thc-backend-go/db/sqlc"
)

type createPackageTagRequest struct {
	TagNumber         string      `json:"tag_number" binding:"required"`
	IsAssigned        bool        `json:"is_assigned"`
	IsProvisional     bool        `json:"is_provisional"`
	IsActive          bool        `json:"is_active"`
	AssignedPackageID nulls.Int64 `json:"assigned_package_id"`
}

func (server *Server) createPackageTag(ctx *gin.Context) {
	var req createPackageTagRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePackageTagParams{
		TagNumber:         req.TagNumber,
		IsAssigned:        req.IsAssigned,
		IsProvisional:     req.IsProvisional,
		IsActive:          req.IsActive,
		AssignedPackageID: req.AssignedPackageID,
	}

	packageTag, err := server.store.CreatePackageTag(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, packageTag)
}

type getPackageTagRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getPackageTag(ctx *gin.Context) {
	var req getPackageTagRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	packageTag, err := server.store.GetPackageTag(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, packageTag)
}

func (server *Server) listPackageTags(ctx *gin.Context) {
	packageTags, err := server.store.ListPackageTags(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, packageTags)
}

type getPackageTagByTagNumberRequest struct {
	TagNumber string `uri:"tag_number" binding:"required"`
}

func (server *Server) getPackageTagByTagNumber(ctx *gin.Context) {
	var req getPackageTagByTagNumberRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	packageTag, err := server.store.GetPackageTagByTagNumber(ctx, req.TagNumber)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, packageTag)
}

func (server *Server) deletePackageTag(ctx *gin.Context) {
	var req getPackageTagRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := server.store.DeletePackageTag(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "package tag deleted"})
}

// TODO: update package tag