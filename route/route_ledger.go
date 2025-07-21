package route

import (
	"github.com/gin-gonic/gin"
	"ledger/auth"
	"ledger/route/handler"
)

var legerPrefix = "/ledger"

func RegisterLedgerGroup(r *gin.RouterGroup) {
	ledgerR := r.Group(legerPrefix, auth.AuthMiddleware())
	//查询功能
	ledgerR.POST("/ledgers", handler.GetLedgerList)
	ledgerR.GET("/:ledger_id/detail", handler.GetLedger)

	//新增功能
	ledgerR.POST("", handler.CreateLedger)

	//编辑功能
	//r.PUT("/:ledger-id", handler.UpdateLedger)
}
