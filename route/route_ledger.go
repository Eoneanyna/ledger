package route

import (
	"github.com/gin-gonic/gin"
	"ledger/auth"
	"ledger/route/handler"
)

var legerPrefix = "/leger"

func RegisterLedgerGroup(r *gin.RouterGroup) {
	r.Group(legerPrefix, auth.AuthMiddleware())
	r.POST("/get/ledger", handler.GetLedgerList)
	//r.POST(legerPrefix, CreateLege)
}
