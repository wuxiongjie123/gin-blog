package api

import (
	"gin-blog/pkg/app"
	"gin-blog/pkg/e"
	"gin-blog/pkg/util"
	"gin-blog/service/auth_service"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	//username := c.Query("username")
	//password := c.Query("password")
	//
	//valid := validation.Validation{}
	//a := auth{Username:username,Password:password}
	//ok, _:= valid.Valid(&a)
	//
	//data := make(map[string]interface{})
	//code := e.INVALID_PARAMS
	//if ok {
	//	isExist := models.CheckAuth(username,password)
	//	if isExist {
	//		token, err := util.GenerateToken(username,password)
	//		if err != nil{
	//			code = e.ERROR_AUTH_TOKEN
	//		}else {
	//			data["token"] = token
	//			code = e.SUCCESS
	//		}
	//	}else {
	//		code = e.ERROR_AUTH
	//	}
	//}else {
	//	for _,err :=range valid.Errors{
	//		logging.Info(err.Key,err.Message)
	//	}
	//}
	//c.JSON(http.StatusOK,gin.H{
	//	"code":code,
	//	"msg":e.GetMsg(code),
	//	"data":data,
	//})

	appG := app.Gin{C: c}
	valid := validation.Validation{}

	username := c.Query("username")
	password := c.Query("password")

	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if ! ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.Auth{Username: username, Password: password}
	isExist, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if ! isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}
	token, err := util.GenerateToken(username, password)
	if err!= nil{
		appG.Response(http.StatusInternalServerError,e.ERROR_AUTH_TOKEN,nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}
