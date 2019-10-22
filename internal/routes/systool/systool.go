package systool
import (
    "os/exec"
    "github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.Engine) {
    r.GET("/api/systool/url/open", urlOpen)
}

func urlOpen(c *gin.Context) {
        /*
        import(
            "os/exec"
        )
        */
        if c.Request.Body == nil {
            c.AbortWithStatus(404)
            return
        }
        params := struct {
            Url    	string `form:"url"`
        }{}
        c.Bind(&params)
                
        cmd := exec.Command("xdg-open", params.Url)
        err := cmd.Start()
        if err!=nil {
            c.JSON(500, gin.H{"err":err})
            return    
        }
        c.JSON(200, gin.H{"cmd":"xdg-open","url":params.Url})        
}

