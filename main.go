package main

import (
	"github.com/marcos9966/rhumdapp/pkg/resmgr"
    "strconv"

    "github.com/gin-gonic/gin"
	"github.com/phayes/freeport"
	"github.com/zserge/webview"
)

var (
	resMgr *resmgr.Resource
	router *gin.Engine
)

func main() {
	resMgr = &resmgr.Resource{
		Root:"app",
	}

	// find first free port to listen with gingonic
	p, err := freeport.GetFreePort()
	if err != nil {
		panic(err)
	}
	port := strconv.Itoa(p)
    router = gin.New()
    router.Use(gin.Logger())
    router.Use(gin.Recovery())
    router.Use(MiddlewareCors())

	// open webview UI
	w := webview.New(webview.Settings{
		Title: "RhumDapp",
        Resizable:true,
        Width:1366,
        Height:720,
        Debug:true,
	})
	w.Eval("window.location = 'http://localhost:"+port+"/index.html';")
	defer w.Exit()
    w.Dispatch(func() {
		go func() {
			handleRoutes(w)
			router.Run("localhost:"+port)
		}()
    })
	w.Run()
}

func MiddlewareCors() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Max-Age", "86400")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
        c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(200)
        } else {
            c.Next()
        }
    }
}
