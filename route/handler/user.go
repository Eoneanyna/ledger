package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"ledger/auth"
	"ledger/database"
	"ledger/my_err"
)

type RegisterUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func RegisterUser(r *gin.Context) {

	var req RegisterUserReq
	err := r.ShouldBindBodyWithJSON(&req)
	if err != nil {
		r.JSON(my_err.ErrInputForm.Code(), gin.H{"error": my_err.ErrInputForm.Error()})
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
		r.JSON(my_err.ErrDataBaseFail.Code(), gin.H{"error": my_err.ErrDataBaseFail.Error()})
		return
	}
	r.JSON(200, gin.H{
		"message": "注册成功",
	})
}

type LoginUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginUser(r *gin.Context) {
	var req LoginUserReq
	err := r.ShouldBindBodyWithJSON(&req)
	if err != nil {
		r.JSON(my_err.ErrUserNotFound.Code(), gin.H{"error": my_err.ErrUserNotFound.Error()})
		return
	}
	//数据库查询用户
	var user database.User
	user, err = database.GetUserByName(req.Username)
	if err != nil {
		log.Errorf("err = %+v, req = %+v", err, req)
		r.JSON(my_err.ErrDataBaseFail.Code(), gin.H{"error": my_err.ErrDataBaseFail.Error()})
		return
	}

	if user.Password != req.Password {
		r.JSON(my_err.ErrInvalidCredentials.Code(), gin.H{"error": my_err.ErrInvalidCredentials.Error()})
		return
	}

	//生成token
	token, err := auth.GenerateToken(fmt.Sprintf("%d", user.Id))
	if err != nil {
		log.Errorf("err = %+v, req = %+v", err, req)
		r.JSON(my_err.ErrDataBaseFail.Code(), gin.H{"error": my_err.ErrDataBaseFail.Error()})
		return
	}
	r.JSON(200, gin.H{
		"token": token,
	})
}
