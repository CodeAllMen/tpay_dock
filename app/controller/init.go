package controller

import (
	"github.com/gin-gonic/gin"
	"tpay_dock/app/libs"
	"tpay_dock/app/models"
)

func InitTables(ctx *gin.Context) {
	models.CreateTable()
	libs.Success(ctx, "create table")
}
