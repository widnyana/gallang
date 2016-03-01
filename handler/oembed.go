package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/widnyana/gallang/handler/processor"
)

// OEmbedService provide logic that will return oEmbed data to you
// as JSON formatted response
func OEmbedService(ctx *gin.Context) {
	url := ctx.DefaultQuery("url", "")
	if len(url) == 0 {
		ctx.String(http.StatusLengthRequired, "Invalid URL.")
	} else {
		svc := processor.OEmbedParser()
		item := svc.FindItem(url)

		if item != nil {
			info, err := item.FetchOembed(url, nil)
			if err != nil {
				ctx.String(http.StatusOK, fmt.Sprintf("An error occured: %s\n", err.Error()))
			} else {
				if info.Status >= 300 {
					ctx.String(http.StatusOK, fmt.Sprintf("Response status code is: %d\n", info.Status))
				} else {
					ctx.JSON(http.StatusOK, info)
				}
			}
		} else {
			ctx.String(http.StatusOK, fmt.Sprintf("oEmbed for URL: %s Not found", url))
		}
	}
}
