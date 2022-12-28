package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/nulls"
	"github.com/shopspring/decimal"
	"net/http"
	db "pixel-thc-backend-go/db/sqlc"
	"time"
)

type createLabTestRequest struct {
	TestName                string          `json:"test_name" binding:"required"`
	BatchCode               string          `json:"batch_code" binding:"required"`
	TestIDCode              string          `json:"test_id_code" binding:"required"`
	LabFacilityName         string          `json:"lab_facility_name" binding:"required"`
	TestPerformedDateTime   time.Time       `json:"test_performed_date_time" binding:"required"`
	OverallPassed           bool            `json:"overall_passed" binding:"required"`
	TestTypeName            string          `json:"test_type_name" binding:"required"`
	TestPassed              bool            `json:"test_passed" binding:"required"`
	TestComment             string          `json:"test_comment" binding:"required"`
	ThcTotalPercent         decimal.Decimal `json:"thc_total_percent" binding:"required"`
	ThcTotalValue           decimal.Decimal `json:"thc_total_value" binding:"required"`
	CbdPercent              decimal.Decimal `json:"cbd_percent" binding:"required"`
	CbdValue                decimal.Decimal `json:"cbd_value" binding:"required"`
	TerpeneTotalPercent     decimal.Decimal `json:"terpene_total_percent" binding:"required"`
	TerpeneTotalValue       decimal.Decimal `json:"terpene_total_value" binding:"required"`
	ThcAPercent             decimal.Decimal `json:"thc_a_percent" binding:"required"`
	ThcAValue               decimal.Decimal `json:"thc_a_value" binding:"required"`
	Delta9ThcPercent        decimal.Decimal `json:"delta_9_thc_percent" binding:"required"`
	Delta9ThcValue          decimal.Decimal `json:"delta_9_thc_value" binding:"required"`
	Delta8ThcPercent        decimal.Decimal `json:"delta_8_thc_percent" binding:"required"`
	Delta8ThcValue          decimal.Decimal `json:"delta_8_thc_value" binding:"required"`
	ThcVPercent             decimal.Decimal `json:"thc_v_percent" binding:"required"`
	ThcVValue               decimal.Decimal `json:"thc_v_value" binding:"required"`
	CbdAPercent             decimal.Decimal `json:"cbd_a_percent" binding:"required"`
	CbdAValue               decimal.Decimal `json:"cbd_a_value" binding:"required"`
	CbnPercent              decimal.Decimal `json:"cbn_percent" binding:"required"`
	CbnValue                decimal.Decimal `json:"cbn_value" binding:"required"`
	CbgAPercent             decimal.Decimal `json:"cbg_a_percent" binding:"required"`
	CbgAValue               decimal.Decimal `json:"cbg_a_value" binding:"required"`
	CbgPercent              decimal.Decimal `json:"cbg_percent" binding:"required"`
	CbgValue                decimal.Decimal `json:"cbg_value" binding:"required"`
	CbcPercent              decimal.Decimal `json:"cbc_percent" binding:"required"`
	CbcValue                decimal.Decimal `json:"cbc_value" binding:"required"`
	TotalCannabinoidPercent decimal.Decimal `json:"total_cannabinoid_percent" binding:"required"`
	TotalCannabinoidValue   decimal.Decimal `json:"total_cannabinoid_value" binding:"required"`
}

func (server *Server) createLabTest(ctx *gin.Context) {
	// Validate input
	var req createLabTestRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create lab test
	arg := db.CreateLabTestParams{
		TestName:                req.TestName,
		BatchCode:               req.BatchCode,
		TestIDCode:              req.TestIDCode,
		LabFacilityName:         req.LabFacilityName,
		TestPerformedDateTime:   req.TestPerformedDateTime,
		OverallPassed:           req.OverallPassed,
		TestTypeName:            req.TestTypeName,
		TestPassed:              req.TestPassed,
		TestComment:             req.TestComment,
		ThcTotalPercent:         req.ThcTotalPercent,
		ThcTotalValue:           req.ThcTotalValue,
		CbdPercent:              req.CbdPercent,
		CbdValue:                req.CbdValue,
		TerpeneTotalPercent:     req.TerpeneTotalPercent,
		TerpeneTotalValue:       req.TerpeneTotalValue,
		ThcAPercent:             req.ThcAPercent,
		ThcAValue:               req.ThcAValue,
		Delta9ThcPercent:        req.Delta9ThcPercent,
		Delta9ThcValue:          req.Delta9ThcValue,
		Delta8ThcPercent:        req.Delta8ThcPercent,
		Delta8ThcValue:          req.Delta8ThcValue,
		ThcVPercent:             req.ThcVPercent,
		ThcVValue:               req.ThcVValue,
		CbdAPercent:             req.CbdAPercent,
		CbdAValue:               req.CbdAValue,
		CbnPercent:              req.CbnPercent,
		CbnValue:                req.CbnValue,
		CbgAPercent:             req.CbgAPercent,
		CbgAValue:               req.CbgAValue,
		CbgPercent:              req.CbgPercent,
		CbgValue:                req.CbgValue,
		CbcPercent:              req.CbcPercent,
		CbcValue:                req.CbcValue,
		TotalCannabinoidPercent: req.TotalCannabinoidPercent,
		TotalCannabinoidValue:   req.TotalCannabinoidValue,
	}

	labTest, err := server.store.CreateLabTest(ctx.Request.Context(), arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, labTest)
}

type getLabTestRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getLabTest(ctx *gin.Context) {
	// Validate input
	var req getLabTestRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get lab test
	labTest, err := server.store.GetLabTest(ctx.Request.Context(), nulls.NewInt64(req.ID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, labTest)
}

func (server *Server) listLabTests(ctx *gin.Context) {
	labTests, err := server.store.ListLabTests(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, labTests)
}

func (server *Server) deleteLabTest(ctx *gin.Context) {
	// Validate input
	var req getLabTestRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Delete lab test
	if err := server.store.DeleteLabTest(ctx.Request.Context(), nulls.NewInt64(req.ID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

// TODO: Add update lab test
