package userhandler

import (
	"github.com/gin-gonic/gin"
	"golang-simple-web-api/component/appctx"
	userstorage "golang-simple-web-api/modules/user/storage"
	"net/http"
)

func ListUser(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := userstorage.ListUser(appctx.GetMainDBConnection())
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, users)
	}
}
