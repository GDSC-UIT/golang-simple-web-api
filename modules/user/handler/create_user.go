package userhandler

import (
	"github.com/gin-gonic/gin"
	"golang-simple-web-api/component/appctx"
	usermodel "golang-simple-web-api/modules/user/model"
	userstorage "golang-simple-web-api/modules/user/storage"
	"net/http"
)

func CreateUser(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqCreateUser usermodel.ReqCreateUser
		if err := c.ShouldBindJSON(&reqCreateUser); err != nil {
			panic(err)
		}

		user := usermodel.User{
			Name: reqCreateUser.Name,
		}
		err := userstorage.CreateUser(appctx.GetMainDBConnection(), &user)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusCreated, user)
	}
}
