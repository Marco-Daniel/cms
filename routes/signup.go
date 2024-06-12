package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignupRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SignupHandler(ctx *gin.Context) {
	var req SignupRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": req.Username, "password": req.Password})
}

func RegisterSignupRoute(r *gin.Engine) {
	r.POST("/signup", SignupHandler)
}
