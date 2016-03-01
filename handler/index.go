package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Hello will return a plain text response.
func Hello(ctx *gin.Context) {
	type resp struct {
		Msg string `json:"message"`
	}

	r := resp{
		Msg: "hello, please check https://godoc.org/github.com/widnyana/oembed for mor info.",
	}

	ctx.JSON(http.StatusOK, r)
}
