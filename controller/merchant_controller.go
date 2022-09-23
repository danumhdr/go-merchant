package controller

import (
	"go-merchant/repository"
	"go-merchant/service/implementation"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var merchanRepository repository.MerchantRepository

func GetMerchantTransaction(context *gin.Context) {
	tokenString := context.GetHeader("Authorization")
	userid, tokenerr := implementation.NewAuthServe().ValidateToken(tokenString)
	if tokenerr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": tokenerr.Error()})
		context.Abort()
		return
	}
	log.Println(userid)
	month, _ := strconv.Atoi(context.Query("month"))
	page, _ := strconv.Atoi(context.Query("page"))
	limit, _ := strconv.Atoi(context.Query("limit"))
	result, err := implementation.NewMerchantServe(merchanRepository).GetReportMerchant(userid, month, limit, page)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "OK", "data": result, "count": len(result), "page": page})
}

func GetOutletTransaction(context *gin.Context) {
	tokenString := context.GetHeader("Authorization")
	userid, _ := implementation.NewAuthServe().ValidateToken(tokenString)
	month, _ := strconv.Atoi(context.Query("month"))
	page, _ := strconv.Atoi(context.Query("page"))
	limit, _ := strconv.Atoi(context.Query("limit"))
	result, err := implementation.NewMerchantServe(merchanRepository).GetReportOutlet(userid, month, limit, page)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "OK", "data": result, "count": len(result), "page": page})
}
