package api

import (
	"github.com/gin-gonic/gin"
	"user/dao"
	"user/service"
)



type RegisterRequest struct {
	Username string
	Email string
	Password string
}

type RegisterResponse struct {
	UserInfo *service.UserInfoDTO `json:"user_info"`
}

func RegisterHandler(c *gin.Context) {

	// 获取post参数
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")


	registerVO := service.RegisterUserVO{
		Username: username,
		Password: password,
		Email:    email,
	}

	userService := service.MakeUserServiceImpl(&dao.UserDAOImpl{})

	userInfoDTO,err := userService.Register(c,&registerVO)


	if err != nil {
		c.JSON(200,gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(200, userInfoDTO)
	}

}


