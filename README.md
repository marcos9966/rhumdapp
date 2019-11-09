# rhumdapp

rhumdapp is based on [webview](https://github.com/zserge/webview) go package.
The purpose of rhumdapp is to be the starting package to write desktop application using go, javascript and html.
Currently the project is tested on linux mint distro, so the following instructions are referred to debian based distro.

### Getting started

Install go ([see documentation](https://golang.org/doc/install)):

create your user github (For example myuser) in folder ~/go/src/github.com

go to user folder (Example myuser)
and download rhumdapp:
```
cd ~/go/src/github.com/myuser
git clone https://github.com/marcos9966/rhumdapp.git
cd rhumdapp
```

replace marcos9966 reference with your repository reference contained in the files
main.go and route.go (for example using vi editor)
```
vi main.go
vi router.go
```


download go-bindata utility (used in this project)
```
go get -u github.com/jteeuwen/go-bindata/...
```

download dependecies
```
go get -u github.com/gin-gonic/gin
go get -u github.com/zserge/webview
go get -u github.com/phayes/freeport
```

If you have problem downloading webview  (Package webkit2gtk-4.0 requested by package webview)
run follow commands
```
sudo apt-get update
sudo apt-get install build-essential libgtk-3-dev
sudo apt-get install libwebkit2gtk-4.0-dev
```

build package
```
cmd/build.sh
```

maybe you may see some notices about deprecated calls, however if it terminated without errors, you can run application with
```
cmd/rhumdapp
```

### Examples

## Adding a new api route

```
cd routes
mkdir example
cd example
vi example.go
```

and insert follow text

```
// Example
package example
import (
    "github.com/gin-gonic/gin"
)

// Rhumdapp Standard function to add package routes
func AddRoutes(r *gin.Engine) {
    r.GET("/api/example/test", test)
}

// test api route
func test(c *gin.Context) {
        if c.Request.Body == nil {
            c.AbortWithStatus(404)
            return
        }
        params := struct {
            Text    	string `form:"text"`
        }{}
        c.Bind(&params)
        if params.Text=="" {
            c.JSON(500, gin.H{"err":"text is mandatory"})
            return
        }
        c.JSON(200, gin.H{
            "result":params.Text,
        })
}
```
change into root folder
```
cd ../..
```
and change router.go
```
vi router.go
```
add import package
add follow code line in import section
```
    "github.com/myuser/rhumdapp/routes/examples"
```
add code after comment
// add here other custom routes
```
    example.AddRoutes(router)
```

## Adding api call in vue component
change first component example 
```
vi res/app/cmp/first/first.vue
```
add html into template html (after h5 html tag)
```
<v-btn rounded color="primary" dark>test</v-btn>
<v-text-field color="success" v-model="text" ></v-text-field>

```
change module exports data (add "text" json tag)
```
module.exports = {
    data: function() {
        return {
            text: ""
        }
    }
```
then add methods section
```
    mounted: function() {
        console.log("first component mounted");
    },
    methods: {
        test: function(){
            var self = this;
            app.api({
                method:"GET",
                url:"/api/example/test",
                params:{text:self.text}
            }).then(data => {
                console.log("api test")
                console.log(data)
            })
            .catch(error => {
                console.log("api test - ERROR")
                console.log(error)
            });
        }
    }


```