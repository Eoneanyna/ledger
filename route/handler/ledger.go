package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"ledger/auth"
	"ledger/database"
	"ledger/my_err"
	"net/http"
)

type GetLedgerListReq struct {
	//日期
	StartTimestamp int64 `json:"start_timestamp" binding:"required"`
	EndTimestamp   int64 `json:"end_timestamp" binding:"required"`
	//分页
	Page     int `utils:"page" binding:"required"` // 当前页码,从1开始
	PageSize int `utils:"pageSize"`                // 每页显示的条目数，默认是10
}

type LedgerList struct {
	Id int64 `json:"id" xorm:"pk autoincr"`
	//金额
	Amount int `json:"amount"`
	//来源 支付宝、微信、银行卡等
	AmountFrom string `json:"amount_from"`
	//时间戳
	Timestamp int64 `json:"timestamp"`
	//描述
	Description string `json:"description"`
}
type GetLedgerListResp struct {
	Ledgers []LedgerList `json:"ledgers"`
	Total   int64        `json:"total"`
}

func GetLedgerList(r *gin.Context) {
	var req GetLedgerListReq
	err := r.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Errorf("err = %+v, req = %+v", err, req)
		r.JSON(my_err.ErrServer.Code(), gin.H{"error": my_err.ErrServer.Error()})
	}

	user_id := r.GetInt64(auth.AuthUserIDKey)

	ledgers, total, err := database.GetLedgerList(user_id, req.StartTimestamp, req.EndTimestamp, req.Page, req.PageSize)
	if err != nil {
		log.Error(fmt.Sprintf("err = %+v, req = %+v", err, req))
		r.JSON(my_err.ErrDataBaseFail.Code(), gin.H{"error": my_err.ErrDataBaseFail.Error()})
		return
	}

	//组装回复
	var resp GetLedgerListResp
	for _, v := range ledgers {
		ledger := LedgerList{
			Id:          v.Id,
			Amount:      v.Amount,
			AmountFrom:  v.AmountFrom,
			Timestamp:   v.Timestamp,
			Description: v.Description,
		}
		resp.Ledgers = append(resp.Ledgers, ledger)
	}
	resp.Total = total

	r.JSON(http.StatusOK, resp)
}
