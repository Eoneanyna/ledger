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
)

type GetTag struct {
	TagId    int64  `json:"tag_id"`
	TagName  string `json:"tag_name"`
	TagTopic int    `json:"tag_topic"`
}
type GetTagListResp struct {
	GetTagList map[int][]GetTag `json:"get_tag_list"`
	Total      int              `json:"total"`
}

func GetTagListHandler(c *gin.Context) {
	userId := c.GetInt64(auth.AuthUserIDKey)

	tags, err := database.GetTagList(userId)
	if err != nil {
		log.Error(fmt.Sprintf("err = %+v", err))
		resp := utils.Resp{
			Code: my_err.ErrDataBaseFail.Code(),
			Msg:  my_err.ErrDataBaseFail.Error(),
			Data: nil,
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	//组装回复
	var resp GetTagListResp
	resp.GetTagList = make(map[int][]GetTag)
	for _, v := range tags {
		ledger := GetTag{
			TagId:    v.TagId,
			TagName:  v.TagName,
			TagTopic: v.TagTopic,
		}
		resp.GetTagList[v.TagTopic] = append(resp.GetTagList[v.TagTopic], ledger)
		resp.Total++
	}

	c.JSON(http.StatusOK, resp)
}

type CreateTagReq struct {
	TagName  string `json:"tag_name" binding:"required"`
	TagTopic int    `json:"tag_topic" binding:"required"`
}

func CreateTagHandler(c *gin.Context) {
	var req CreateTagReq
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Errorf("err = %+v, req = %+v", err, req)
		resp := utils.Resp{
			Code: my_err.ErrServer.Code(),
			Msg:  my_err.ErrServer.Error(),
			Data: nil,
		}
		c.JSON(http.StatusOK, resp)
	}

	//查找当前用户已有标签数量
	tagCount, err := database.GetTagCount(c.GetInt64(auth.AuthUserIDKey))
	if err != nil {
		log.Error(fmt.Sprintf("err = %+v", err))
		resp := utils.Resp{
			Code: my_err.ErrDataBaseFail.Code(),
			Msg:  my_err.ErrDataBaseFail.Error(),
			Data: nil,
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	if tagCount >= 10 {
		resp := utils.Resp{
			Code: my_err.ErrTagLimit.Code(),
			Msg:  my_err.ErrTagLimit.Error(),
			Data: nil,
		}
		c.JSON(http.StatusOK, resp)
	}

	tag := database.LedgerTag{
		TagName:  req.TagName,
		TagTopic: req.TagTopic,
		UserId:   c.GetInt64(auth.AuthUserIDKey),
	}
	err = database.CreateTag(&tag)
	if err != nil {
		log.Error(fmt.Sprintf("err = %+v", err))
		resp := utils.Resp{
			Code: my_err.ErrDataBaseFail.Code(),
			Msg:  my_err.ErrDataBaseFail.Error(),
			Data: nil,
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	resp := utils.Resp{
		Code: 200,
		Msg:  "添加成功",
		Data: nil,
	}

	c.JSON(http.StatusOK, resp)
}

type UpdateTagReq struct {
	TagId    int64  `json:"tag_id" binding:"required"`
	TagName  string `json:"tag_name"`
	TagTopic int    `json:"tag_topic"`
}

func UpdateTagHandler(c *gin.Context) {
	var req UpdateTagReq
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Errorf("err = %+v, req = %+v", err, req)
		resp := utils.Resp{
			Code: my_err.ErrServer.Code(),
			Msg:  my_err.ErrServer.Error(),
			Data: nil,
		}
		c.JSON(http.StatusOK, resp)
	}

	//userId := c.GetInt64(auth.AuthUserIDKey)

	Update := database.LedgerTag{
		TagName:  req.TagName,
		TagTopic: req.TagTopic,
	}
	err = database.UpdateLedgerTag(req.TagId, Update)
	if err != nil {
		resp := utils.Resp{
			Code: my_err.ErrDataBaseFail.Code(),
			Msg:  my_err.ErrDataBaseFail.Error(),
			Data: nil,
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	resp := utils.Resp{
		Code: 200,
		Msg:  "更新成功",
		Data: nil,
	}
	c.JSON(http.StatusOK, resp)
}
