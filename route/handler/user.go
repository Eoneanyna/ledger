package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"ledger/auth"
	"ledger/database"
	"ledger/my_err"
	"ledger/utils"
)

type RegisterUserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

func RegisterUser(c *gin.Context) {

	var req RegisterUserReq
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		resp := utils.Resp{
			Code: my_err.ErrInputForm.Code(),
			Msg:  my_err.ErrInputForm.Error(),
			Data: nil,
		}
		c.JSON(my_err.ErrInputForm.Code(), resp)
		return
	}

	var user = database.User{
		Name:     req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: req.Password,
	}
	err = database.InsertUser(user)
	if err != nil {
		log.Errorf("err = %+v, req = %+v", err, req)
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
		Msg:  "注册成功",
		Data: nil,
	}
	c.JSON(200, resp)
}

type LoginUserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserResp struct {
	Token string `json:"token"`
}

func LoginUser(c *gin.Context) {
	var req LoginUserReq
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Errorf("err = %+v, req = %+v", err, req)
		resp := utils.Resp{
			Code: my_err.ErrUserNotFound.Code(),
			Msg:  my_err.ErrUserNotFound.Error(),
			Data: nil,
		}
		c.JSON(my_err.ErrUserNotFound.Code(), resp)
		return
	}
	//数据库查询用户
	var user database.User
	user, err = database.GetUserByName(req.Username)
	if err != nil {
		log.Errorf("err = %+v, req = %+v", err, req)
		resp := utils.Resp{
			Code: my_err.ErrDataBaseFail.Code(),
			Msg:  my_err.ErrDataBaseFail.Error(),
			Data: nil,
		}
		c.JSON(my_err.ErrDataBaseFail.Code(), resp)
		return
	}

	if user.Password != req.Password {
		log.Debugf("err = %+v, req = %+v", err, req)
		resp := utils.Resp{
			Code: my_err.ErrInvalidCredentials.Code(),
			Msg:  my_err.ErrInvalidCredentials.Error(),
			Data: nil,
		}
		c.JSON(my_err.ErrInvalidCredentials.Code(), resp)
		return
	}

	//生成token
	token, err := auth.GenerateToken(user.Id)
	if err != nil {
		log.Errorf("err = %+v, req = %+v", err, req)
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
		Msg:  "登录成功",
		Data: LoginUserResp{
			Token: token,
		},
	}
	c.JSON(200, resp)
}
