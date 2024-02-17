package testapi

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ClientCreateEvent(ctx *gin.Context) {
	var req Params
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("ClientCreateEvent: ", "ctx.ShouldBindJSON(&req) err")
		//loggerwrite.Logger().WithFields(logrus.Fields{
		//	"name": "ctx.ShouldBindJSON(&req)",
		//}).Error("err:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "")
}
