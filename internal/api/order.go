package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/nulls"
	"github.com/shopspring/decimal"
	"net/http"
	"pixel-thc-backend-go/internal/db/sqlc"
	"time"
)

type createOrderRequest struct {
	ScheduledPackDateTime     time.Time       `json:"scheduled_pack_date_time"`
	ScheduledShipDateTime     time.Time       `json:"scheduled_ship_date_time"`
	ScheduledDeliveryDateTime time.Time       `json:"scheduled_delivery_date_time"`
	ActualPackDateTime        nulls.Time      `json:"actual_pack_date_time"`
	ActualShipDateTime        nulls.Time      `json:"actual_ship_date_time"`
	ActualDeliveryDateTime    nulls.Time      `json:"actual_delivery_date_time"`
	OrderTotal                decimal.Decimal `json:"order_total"`
	Notes                     string          `json:"notes"`
	Status                    string          `json:"status"`
	CustomerName              string          `json:"customer_name"`
}

func (server *Server) createOrder(ctx *gin.Context) {
	var req createOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	arg := db.CreateOrderParams{
		ScheduledPackDateTime:     req.ScheduledPackDateTime,
		ScheduledShipDateTime:     req.ScheduledShipDateTime,
		ScheduledDeliveryDateTime: req.ScheduledDeliveryDateTime,
		ActualPackDateTime:        req.ActualPackDateTime,
		ActualShipDateTime:        req.ActualShipDateTime,
		ActualDeliveryDateTime:    req.ActualDeliveryDateTime,
		OrderTotal:                req.OrderTotal,
		Notes:                     req.Notes,
		Status:                    req.Status,
		CustomerName:              req.CustomerName,
	}

	order, err := server.store.CreateOrder(ctx.Request.Context(), arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, order)
}

type getOrderRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getOrder(ctx *gin.Context) {
	var req getOrderRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	order, err := server.store.GetOrder(ctx, req.ID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, order)
}

func (server *Server) listOrders(ctx *gin.Context) {
	orders, err := server.store.ListOrders(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, orders)
}

type updateOrderRequest struct {
	ID                        int64               `uri:"id" binding:"required,min=1"`
	ScheduledPackDateTime     nulls.Time          `json:"scheduled_pack_date_time"`
	ScheduledShipDateTime     nulls.Time          `json:"scheduled_ship_date_time"`
	ScheduledDeliveryDateTime nulls.Time          `json:"scheduled_delivery_date_time"`
	ActualPackDateTime        nulls.Time          `json:"actual_pack_date_time"`
	ActualShipDateTime        nulls.Time          `json:"actual_ship_date_time"`
	ActualDeliveryDateTime    nulls.Time          `json:"actual_delivery_date_time"`
	OrderTotal                decimal.NullDecimal `json:"order_total"`
	Notes                     nulls.String        `json:"notes"`
	Status                    nulls.String        `json:"status"`
	CustomerName              nulls.String        `json:"customer_name"`
}

func (server *Server) updateOrder(ctx *gin.Context) {
	var req updateOrderRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	arg := db.UpdateOrderParams{
		ID:                        req.ID,
		ScheduledPackDateTime:     req.ScheduledPackDateTime,
		ScheduledShipDateTime:     req.ScheduledShipDateTime,
		ScheduledDeliveryDateTime: req.ScheduledDeliveryDateTime,
		ActualPackDateTime:        req.ActualPackDateTime,
		ActualShipDateTime:        req.ActualShipDateTime,
		ActualDeliveryDateTime:    req.ActualDeliveryDateTime,
		OrderTotal:                req.OrderTotal,
		Notes:                     req.Notes,
		Status:                    req.Status,
		CustomerName:              req.CustomerName,
	}

	order, err := server.store.UpdateOrder(ctx, arg)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, order)
}

type deleteOrderRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteOrder(ctx *gin.Context) {
	var req deleteOrderRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := server.store.DeleteOrder(ctx, req.ID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, nil)
}
