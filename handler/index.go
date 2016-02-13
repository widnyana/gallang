package handler

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func Hello(ctx *gin.Context) {
    ctx.String(http.StatusOK, "Hello")
}
