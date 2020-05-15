/**
  create by yy on 2019-08-23
*/

package controller

import (
	"github.com/gin-gonic/gin"
	"tpay_dock/app/libs"
)

func Index(ctx *gin.Context) {
	libs.Success(ctx, "index")
}
