package controller

import (
	"go-merchant/model"
	"go-merchant/service/implementation"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignIn(context *gin.Context) {
	var request model.SignInRequest
	var user model.User
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	user, validate := implementation.NewAuthServe().CheckCredential(request.UserName, request.Password)
	if validate != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": validate.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "OK", "data": user})
}
