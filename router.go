package main

import (
	"errors"
	"net/http"
    "os"
    "path/filepath"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/marcos9966/rhumdapp/routes/systool"
	"github.com/zserge/webview"
)

func handleRoutes(w webview.WebView) {

    router.GET("/api/action/quit", func(c *gin.Context) {
        w.Terminate()
        os.Exit(0)
    })

    // add here other custom routes
    systool.AddRoutes(router)

	// if not mapped as route try to open resource file
    router.NoRoute(func(c *gin.Context) {
        var contentType string
        var data []byte
        err:= errors.New("")
        data, err = resMgr.Get(c.Request.URL.String())
        if err != nil {
			c.AbortWithStatus(404)
			return
        }

        contentType = http.DetectContentType(data)
        switch strings.ToLower(filepath.Ext(c.Request.URL.String())) {
            case ".js":
                contentType = "application/x-javascript; charset=utf-8"
            case ".css":
                contentType = "text/css; charset=utf-8"
            case ".html":
                contentType = "text/html; charset=utf-8"
        }
        c.Data(http.StatusOK, contentType, data)
    })
}
