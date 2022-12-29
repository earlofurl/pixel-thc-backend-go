package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/nulls"
	"github.com/shopspring/decimal"
	"net/http"
	db "pixel-thc-backend-go/db/sqlc"
	"time"
)

type createPackageRequest struct {
	SourcePackageID                   nulls.Int64     `json:"source_package_id"`
	TagID                             nulls.Int64     `json:"tag_id"`
	PackageType                       string          `json:"package_type"`
	IsActive                          bool            `json:"is_active"`
	Quantity                          decimal.Decimal `json:"quantity"`
	Notes                             string          `json:"notes"`
	PackagedDateTime                  time.Time       `json:"packaged_date_time"`
	HarvestDateTime                   nulls.Time      `json:"harvest_date_time"`
	LabTestingState                   string          `json:"lab_testing_state"`
	LabTestingStateDateTime           nulls.Time      `json:"lab_testing_state_date_time"`
	IsTradeSample                     bool            `json:"is_trade_sample"`
	IsTestingSample                   bool            `json:"is_testing_sample"`
	ProductRequiresRemediation        bool            `json:"product_requires_remediation"`
	ContainsRemediatedProduct         bool            `json:"contains_remediated_product"`
	RemediationDateTime               nulls.Time      `json:"remediation_date_time"`
	ReceivedDateTime                  nulls.Time      `json:"received_date_time"`
	ReceivedFromManifestNumber        nulls.String    `json:"received_from_manifest_number"`
	ReceivedFromFacilityLicenseNumber nulls.String    `json:"received_from_facility_license_number"`
	ReceivedFromFacilityName          nulls.String    `json:"received_from_facility_name"`
	IsOnHold                          bool            `json:"is_on_hold"`
	ArchivedDate                      nulls.Time      `json:"archived_date"`
	FinishedDate                      nulls.Time      `json:"finished_date"`
	ItemID                            nulls.Int64     `json:"item_id"`
	ProvisionalLabel                  nulls.String    `json:"provisional_label"`
	IsProvisional                     bool            `json:"is_provisional"`
	IsSold                            bool            `json:"is_sold"`
	PpuDefault                        decimal.Decimal `json:"ppu_default"`
	PpuOnOrder                        decimal.Decimal `json:"ppu_on_order"`
	TotalPackagePriceOnOrder          decimal.Decimal `json:"total_package_price_on_order"`
	PpuSoldPrice                      decimal.Decimal `json:"ppu_sold_price"`
	TotalSoldPrice                    decimal.Decimal `json:"total_sold_price"`
	PackagingSuppliesConsumed         bool            `json:"packaging_supplies_consumed"`
	IsLineItem                        bool            `json:"is_line_item"`
	OrderID                           nulls.Int64     `json:"order_id"`
	UomID                             int64           `json:"uom_id"`
	LabTestID                         nulls.Int64     `json:"lab_test_id"`
	FacilityLocationID                int64           `json:"facility_location_id"`
}

func (server *Server) createPackage(ctx *gin.Context) {
	var req createPackageRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	packageParams := db.CreatePackageParams{
		TagID:                             req.TagID,
		PackageType:                       req.PackageType,
		IsActive:                          req.IsActive,
		Quantity:                          req.Quantity,
		Notes:                             req.Notes,
		PackagedDateTime:                  req.PackagedDateTime,
		HarvestDateTime:                   req.HarvestDateTime,
		LabTestingState:                   req.LabTestingState,
		LabTestingStateDateTime:           req.LabTestingStateDateTime,
		IsTradeSample:                     req.IsTradeSample,
		IsTestingSample:                   req.IsTestingSample,
		ProductRequiresRemediation:        req.ProductRequiresRemediation,
		ContainsRemediatedProduct:         req.ContainsRemediatedProduct,
		RemediationDateTime:               req.RemediationDateTime,
		ReceivedDateTime:                  req.ReceivedDateTime,
		ReceivedFromManifestNumber:        req.ReceivedFromManifestNumber,
		ReceivedFromFacilityLicenseNumber: req.ReceivedFromFacilityLicenseNumber,
		ReceivedFromFacilityName:          req.ReceivedFromFacilityName,
		IsOnHold:                          req.IsOnHold,
		ArchivedDate:                      req.ArchivedDate,
		FinishedDate:                      req.FinishedDate,
		ItemID:                            req.ItemID,
		ProvisionalLabel:                  req.ProvisionalLabel,
		IsProvisional:                     req.IsProvisional,
		IsSold:                            req.IsSold,
		PpuDefault:                        req.PpuDefault,
		PpuOnOrder:                        req.PpuOnOrder,
		TotalPackagePriceOnOrder:          req.TotalPackagePriceOnOrder,
		PpuSoldPrice:                      req.PpuSoldPrice,
		TotalSoldPrice:                    req.TotalSoldPrice,
		PackagingSuppliesConsumed:         req.PackagingSuppliesConsumed,
		IsLineItem:                        req.IsLineItem,
		OrderID:                           req.OrderID,
		UomID:                             req.UomID,
		FacilityLocationID:                req.FacilityLocationID,
	}

	arg := db.CreatePackageTxParams{
		SourcePackageID:     req.SourcePackageID,
		CreatePackageParams: packageParams,
	}

	productPackage, err := server.store.CreatePackageTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, productPackage)
}

type getPackageRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getPackage(ctx *gin.Context) {
	var req getPackageRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	productPackage, err := server.store.GetPackage(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, productPackage)
}

type getPackageByTagIDRequest struct {
	TagID int64 `uri:"tag_id" binding:"required,min=1"`
}

func (server *Server) getPackageByTagID(ctx *gin.Context) {
	var req getPackageByTagIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	productPackage, err := server.store.GetPackageByTagID(ctx, nulls.NewInt64(req.TagID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, productPackage)
}

func (server *Server) listPackages(ctx *gin.Context) {
	productPackages, err := server.store.ListPackages(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, productPackages)
}

func (server *Server) deletePackage(ctx *gin.Context) {
	var req getPackageRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := server.store.DeletePackage(ctx, req.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "package deleted"})
}

func (server *Server) listActivePackages(ctx *gin.Context) {
	productPackages, err := server.store.ListActivePackages(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, productPackages)
}

type updatePackageRequest struct {
	ID                                int64               `json:"id" binding:"required,min=1"`
	TagID                             nulls.Int64         `json:"tag_id"`
	PackageType                       nulls.String        `json:"package_type"`
	Quantity                          decimal.NullDecimal `json:"quantity"`
	Notes                             nulls.String        `json:"notes"`
	PackagedDateTime                  nulls.Time          `json:"packaged_date_time"`
	HarvestDateTime                   nulls.Time          `json:"harvest_date_time"`
	LabTestingState                   nulls.String        `json:"lab_testing_state"`
	LabTestingStateDateTime           nulls.Time          `json:"lab_testing_state_date_time"`
	IsTradeSample                     nulls.Bool          `json:"is_trade_sample"`
	IsTestingSample                   nulls.Bool          `json:"is_testing_sample"`
	ProductRequiresRemediation        nulls.Bool          `json:"product_requires_remediation"`
	ContainsRemediatedProduct         nulls.Bool          `json:"contains_remediated_product"`
	RemediationDateTime               nulls.Time          `json:"remediation_date_time"`
	ReceivedDateTime                  nulls.Time          `json:"received_date_time"`
	ReceivedFromManifestNumber        nulls.String        `json:"received_from_manifest_number"`
	ReceivedFromFacilityLicenseNumber nulls.String        `json:"received_from_facility_license_number"`
	ReceivedFromFacilityName          nulls.String        `json:"received_from_facility_name"`
	IsOnHold                          nulls.Bool          `json:"is_on_hold"`
	ArchivedDate                      nulls.Time          `json:"archived_date"`
	FinishedDate                      nulls.Time          `json:"finished_date"`
	ItemID                            nulls.Int64         `json:"item_id"`
	ProvisionalLabel                  nulls.String        `json:"provisional_label"`
	IsProvisional                     nulls.Bool          `json:"is_provisional"`
	IsSold                            nulls.Bool          `json:"is_sold"`
	PpuDefault                        decimal.NullDecimal `json:"ppu_default"`
	PpuOnOrder                        decimal.NullDecimal `json:"ppu_on_order"`
	TotalPackagePriceOnOrder          decimal.NullDecimal `json:"total_package_price_on_order"`
	PpuSoldPrice                      decimal.NullDecimal `json:"ppu_sold_price"`
	TotalSoldPrice                    decimal.NullDecimal `json:"total_sold_price"`
	PackagingSuppliesConsumed         nulls.Bool          `json:"packaging_supplies_consumed"`
	IsLineItem                        nulls.Bool          `json:"is_line_item"`
	OrderID                           nulls.Int64         `json:"order_id"`
	UomID                             nulls.Int64         `json:"uom_id"`
	IsActive                          nulls.Bool          `json:"is_active"`
}

func (server *Server) updatePackage(ctx *gin.Context) {
	var req updatePackageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	productPackage, err := server.store.UpdatePackage(ctx, db.UpdatePackageParams{
		ID:                                req.ID,
		TagID:                             req.TagID,
		PackageType:                       req.PackageType,
		Quantity:                          req.Quantity,
		Notes:                             req.Notes,
		PackagedDateTime:                  req.PackagedDateTime,
		HarvestDateTime:                   req.HarvestDateTime,
		LabTestingState:                   req.LabTestingState,
		LabTestingStateDateTime:           req.LabTestingStateDateTime,
		IsTradeSample:                     req.IsTradeSample,
		IsTestingSample:                   req.IsTestingSample,
		ProductRequiresRemediation:        req.ProductRequiresRemediation,
		ContainsRemediatedProduct:         req.ContainsRemediatedProduct,
		RemediationDateTime:               req.RemediationDateTime,
		ReceivedDateTime:                  req.ReceivedDateTime,
		ReceivedFromManifestNumber:        req.ReceivedFromManifestNumber,
		ReceivedFromFacilityLicenseNumber: req.ReceivedFromFacilityLicenseNumber,
		ReceivedFromFacilityName:          req.ReceivedFromFacilityName,
		IsOnHold:                          req.IsOnHold,
		ArchivedDate:                      req.ArchivedDate,
		FinishedDate:                      req.FinishedDate,
		ItemID:                            req.ItemID,
		ProvisionalLabel:                  req.ProvisionalLabel,
		IsProvisional:                     req.IsProvisional,
		IsSold:                            req.IsSold,
		PpuDefault:                        req.PpuDefault,
		PpuOnOrder:                        req.PpuOnOrder,
		TotalPackagePriceOnOrder:          req.TotalPackagePriceOnOrder,
		PpuSoldPrice:                      req.PpuSoldPrice,
		TotalSoldPrice:                    req.TotalSoldPrice,
		PackagingSuppliesConsumed:         req.PackagingSuppliesConsumed,
		IsLineItem:                        req.IsLineItem,
		OrderID:                           req.OrderID,
		UomID:                             req.UomID,
		IsActive:                          req.IsActive,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, productPackage)
}
