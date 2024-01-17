package userhandler

import (
	"github.com/gin-gonic/gin"
	"golang-simple-web-api/component/appctx"
	usermodel "golang-simple-web-api/modules/user/model"
	userstorage "golang-simple-web-api/modules/user/storage"
)

func UpdateUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqUpdateUser usermodel.ReqUpdateUser
		if err := c.ShouldBindJSON(&reqUpdateUser); err != nil {
			panic(err)
		}

		userId := c.Param("id")

		userstorage.UpdateUser(appCtx.GetMainDBConnection(), userId, &reqUpdateUser)
	}
}
