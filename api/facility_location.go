package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "pixel-thc-backend-go/db/sqlc"
)

type createFacilityLocationRequest struct {
	FacilityID int64  `json:"facility_id" binding:"required"`
	Name       string `json:"name" binding:"required"`
}

func (server *Server) createFacilityLocation(ctx *gin.Context) {
	var req createFacilityLocationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateFacilityLocationParams{
		FacilityID: req.FacilityID,
		Name:       req.Name,
	}

	facilityLocation, err := server.store.CreateFacilityLocation(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, facilityLocation)
}

type getFacilityLocationRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getFacilityLocation(ctx *gin.Context) {
	var req getFacilityLocationRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	facilityLocation, err := server.store.GetFacilityLocation(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, facilityLocation)
}

func (server *Server) listFacilityLocations(ctx *gin.Context) {
	facilityLocations, err := server.store.ListFacilityLocations(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, facilityLocations)
}

type updateFacilityLocationRequest struct {
	ID         int64  `uri:"id" binding:"required,min=1"`
	FacilityID int64  `json:"facility_id" binding:"required"`
	Name       string `json:"name" binding:"required"`
}

func (server *Server) updateFacilityLocation(ctx *gin.Context) {
	var req updateFacilityLocationRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateFacilityLocationParams{
		ID:         req.ID,
		FacilityID: req.FacilityID,
		Name:       req.Name,
	}

	facilityLocation, err := server.store.UpdateFacilityLocation(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, facilityLocation)
}

type deleteFacilityLocationRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteFacilityLocation(ctx *gin.Context) {
	var req deleteFacilityLocationRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteFacilityLocation(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
