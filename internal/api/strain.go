package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/nulls"
	"github.com/shopspring/decimal"
	"net/http"
	"pixel-thc-backend-go/internal/db/sqlc"
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

type updateStrainRequest struct {
	ID                      int64               `json:"id" binding:"required,min=1"`
	Name                    nulls.String        `json:"name"`
	Type                    nulls.String        `json:"type"`
	YieldAverage            decimal.NullDecimal `json:"yield_average"`
	TerpAverageTotal        decimal.NullDecimal `json:"terp_average_total"`
	Terp1                   nulls.String        `json:"terp_1"`
	Terp1Value              decimal.NullDecimal `json:"terp_1_value"`
	Terp2                   nulls.String        `json:"terp_2"`
	Terp2Value              decimal.NullDecimal `json:"terp_2_value"`
	Terp3                   nulls.String        `json:"terp_3"`
	Terp3Value              decimal.NullDecimal `json:"terp_3_value"`
	Terp4                   nulls.String        `json:"terp_4"`
	Terp4Value              decimal.NullDecimal `json:"terp_4_value"`
	Terp5                   nulls.String        `json:"terp_5"`
	Terp5Value              decimal.NullDecimal `json:"terp_5_value"`
	ThcAverage              decimal.NullDecimal `json:"thc_average"`
	TotalCannabinoidAverage decimal.NullDecimal `json:"total_cannabinoid_average"`
	LightDep2022            nulls.String        `json:"light_dep_2022"`
	FallHarvest2022         nulls.String        `json:"fall_harvest_2022"`
	QuantityAvailable       decimal.NullDecimal `json:"quantity_available"`
}

func (server *Server) updateStrain(ctx *gin.Context) {
	var req updateStrainRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateStrainParams{
		ID:                      req.ID,
		Name:                    req.Name,
		Type:                    req.Type,
		YieldAverage:            req.YieldAverage,
		TerpAverageTotal:        req.TerpAverageTotal,
		Terp1:                   req.Terp1,
		Terp1Value:              req.Terp1Value,
		Terp2:                   req.Terp2,
		Terp2Value:              req.Terp2Value,
		Terp3:                   req.Terp3,
		Terp3Value:              req.Terp3Value,
		Terp4:                   req.Terp4,
		Terp4Value:              req.Terp4Value,
		Terp5:                   req.Terp5,
		Terp5Value:              req.Terp5Value,
		ThcAverage:              req.ThcAverage,
		TotalCannabinoidAverage: req.TotalCannabinoidAverage,
		LightDep2022:            req.LightDep2022,
		FallHarvest2022:         req.FallHarvest2022,
		QuantityAvailable:       req.QuantityAvailable,
	}

	strain, err := server.store.UpdateStrain(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, strain)
}
