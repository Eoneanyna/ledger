package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"ledger/auth"
	"ledger/database"
	"ledger/my_err"
	"ledger/utils"
	"net/http"
	"strconv"
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

func GetLedgerList(c *gin.Context) {
	var req GetLedgerListReq
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Errorf("err = %+v, req = %+v", err, req)
		resp := utils.Resp{
			Code: my_err.ErrServer.Code(),
			Msg:  my_err.ErrServer.Error(),
			Data: nil,
		}
		c.JSON(my_err.ErrServer.Code(), resp)
	}

	userId := c.GetInt64(auth.AuthUserIDKey)

	ledgers, total, err := database.GetLedgerList(userId, req.StartTimestamp, req.EndTimestamp, req.Page, req.PageSize)
	if err != nil {
		log.Error(fmt.Sprintf("err = %+v, req = %+v", err, req))
		resp := utils.Resp{
			Code: my_err.ErrDataBaseFail.Code(),
			Msg:  my_err.ErrDataBaseFail.Error(),
			Data: nil,
		}
		c.JSON(my_err.ErrDataBaseFail.Code(), resp)
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

	c.JSON(http.StatusOK, resp)
}

type CreateLedgerReq struct {
	//金额
	Amount int `json:"amount" binding:"required"`
	//来源 支付宝、微信、银行卡等
	AmountFrom string `json:"amount_from" binding:"required"`
	//时间戳
	Timestamp int64 `json:"timestamp" binding:"required"`
	//描述
	Description string `json:"description" binding:"required"`
}

func CreateLedger(c *gin.Context) {
	var req CreateLedgerReq
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Errorf("err = %+v, req = %+v", err, req)
		resp := utils.Resp{
			Code: my_err.ErrServer.Code(),
			Msg:  my_err.ErrServer.Error(),
			Data: nil,
		}
		c.JSON(my_err.ErrServer.Code(), resp)
	}

	userId := c.GetInt64(auth.AuthUserIDKey)

	err = database.InsertLedger(&database.Ledger{
		UserId:      userId,
		Amount:      req.Amount,
		AmountFrom:  req.AmountFrom,
		Timestamp:   req.Timestamp,
		Description: req.Description,
	})
	if err != nil {
		log.Error(fmt.Sprintf("err = %+v, req = %+v", err, req))
		resp := utils.Resp{
			Code: my_err.ErrDataBaseFail.Code(),
			Msg:  my_err.ErrDataBaseFail.Error(),
			Data: nil,
		}
		c.JSON(my_err.ErrDataBaseFail.Code(), resp)
		return
	}

	resp := utils.Resp{
		Code: 200,
		Msg:  "添加成功",
		Data: nil,
	}
	c.JSON(http.StatusOK, resp)
}

type GetLedgerResp struct {
	//金额
	Amount int `json:"amount" binding:"required"`
	//来源 支付宝、微信、银行卡等
	AmountFrom string `json:"amount_from" binding:"required"`
	//TODO 标签名称
	//时间戳
	Timestamp int64 `json:"timestamp" binding:"required"`
	//描述
	Description string `json:"description" binding:"required"`
}

func GetLedger(c *gin.Context) {
	ledgerIdStr := c.Param("ledger_id")
	ledgerId, err := strconv.ParseInt(ledgerIdStr, 10, 64)
	if err != nil {
		log.Errorf("invalid ledger_id: %v", err)
		resp := utils.Resp{
			Code: my_err.ErrInvalidParam.Code(),
			Msg:  "invalid ledger_id",
			Data: nil,
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	userId := c.GetInt64(auth.AuthUserIDKey)
	findData := &database.Ledger{
		Id:     ledgerId,
		UserId: userId,
	}
	err = database.FindLedger(findData)
	if err != nil {
		log.Errorf("err = %+v, req = %+v", err, ledgerId)
		resp := utils.Resp{
			Code: my_err.ErrDataBaseFail.Code(),
			Msg:  my_err.ErrDataBaseFail.Error(),
			Data: nil,
		}
		c.JSON(my_err.ErrDataBaseFail.Code(), resp)
		return
	}

	resp := utils.Resp{
		Code: 200,
		Msg:  "查询成功",
		Data: GetLedgerResp{
			Amount:      findData.Amount,
			AmountFrom:  findData.AmountFrom,
			Timestamp:   findData.Timestamp,
			Description: findData.Description,
		},
	}
	c.JSON(http.StatusOK, resp)
}
