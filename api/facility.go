package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/nulls"
	"net/http"
	db "pixel-thc-backend-go/db/sqlc"
)

type createFacilityRequest struct {
	Name          string `json:"name" binding:"required"`
	LicenseNumber string `json:"license_number" binding:"required"`
}

func (server *Server) createFacility(ctx *gin.Context) {
	var req createFacilityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateFacilityParams{
		Name:          req.Name,
		LicenseNumber: req.LicenseNumber,
	}

	facility, err := server.store.CreateFacility(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, facility)
}

type getFacilityRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getFacility(ctx *gin.Context) {
	var req getFacilityRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	facility, err := server.store.GetFacility(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, facility)
}

func (server *Server) listFacilities(ctx *gin.Context) {
	facilities, err := server.store.ListFacilities(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, facilities)
}

type updateFacilityRequest struct {
	ID            int64  `uri:"id" binding:"required,min=1"`
	Name          string `json:"name" binding:"required"`
	LicenseNumber string `json:"license_number" binding:"required"`
}

func (server *Server) updateFacility(ctx *gin.Context) {
	var req updateFacilityRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var updateReq createFacilityRequest
	if err := ctx.ShouldBindJSON(&updateReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateFacilityParams{
		ID:            req.ID,
		Name:          nulls.NewString(updateReq.Name),
		LicenseNumber: nulls.NewString(updateReq.LicenseNumber),
	}

	facility, err := server.store.UpdateFacility(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, facility)
}

func (server *Server) deleteFacility(ctx *gin.Context) {
	var req getFacilityRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteFacility(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "facility deleted"})
}
