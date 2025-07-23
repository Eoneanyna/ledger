package route

import (
	"github.com/gin-gonic/gin"
	"ledger/auth"
	"ledger/route/handler"
)

var tagPrefix = "/tag"

func RegisterTagGroup(r *gin.RouterGroup) {
	tagR := r.Group(legerPrefix, auth.AuthMiddleware())
	//查询功能
	tagR.POST("/tags", handler.GetTagListHandler)

	// TODO 新增功能
	//tagR.POST("", handler.CreateTagHandler)

	//TODO 编辑功能
	//更改层级，更改名称
	//tagR.PUT("/tag", handler.UpdateTagHandler)
}
