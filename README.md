# rhumdapp

rhumdapp is based on [webview](https://github.com/zserge/webview) go library.
The purpose of rhumdapp is to be the skeleton to write desktop application using go, javascript and html.
Currently the project is tested on linux mint distro, so the following instruction are referred to debian based distro.

### Getting started

Install go ([see documentation](https://golang.org/doc/install)):

and then make your repository (For example myrepo)

change repo directory (Example myrepo)
and download rhumdapp:
```
cd ~/go/src/github.com/myrepo
git clone https://github.com/marcos9966/rhumdapp.git
```

replace marcos9966 reference with your report reference contained in file
main.go and route.go


download go-bindata utility (used in this project)
```
go get -u github.com/jteeuwen/go-bindata/...
```

download dependecies
```
go get -u go get -u github.com/gin-gonic/gin
go get github.com/zserge/webview
go get github.com/phayes/freeport
```

try to build package
```
cmd/build.sh
```

If you have problem with GTK3 install or Package webkit2gtk-4.0 (requested by package webview)
```
sudo apt-get update
sudo apt-get install build-essential libgtk-3-dev
sudo apt-get install libwebkit2gtk-4.0-dev
```

trying again to build project, maybe you may see some notice, however if compiling without errors, you can run application with
```
cmd/rhumdapp
```
