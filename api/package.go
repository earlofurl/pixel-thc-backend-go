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
	TagID                             nulls.Int64     `json:"tag_id"`
	PackageType                       string          `json:"package_type"`
	Quantity                          decimal.Decimal `json:"quantity"`
	Notes                             nulls.String    `json:"notes"`
	PackagedDateTime                  time.Time       `json:"packaged_date_time"`
	HarvestDateTime                   nulls.Time      `json:"harvest_date_time"`
	LabTestingState                   nulls.String    `json:"lab_testing_state"`
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
	UomID                             nulls.Int64     `json:"uom_id"`
}

func (server *Server) createPackage(ctx *gin.Context) {
	var req createPackageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePackageParams{
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
	}

	productPackage, err := server.store.CreatePackage(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, productPackage)
}
