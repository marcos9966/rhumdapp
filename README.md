# rhumdapp

rhumdapp is based on [webview](https://github.com/zserge/webview) go library.
The purpose of rhumdapp is to be the skeleton to write desktop application using go, javascript and html.
Currently the project is tested on linux mint distro, so the following instruction are referred to debian based distro.

### Getting started

Install go ([see documentation](https://golang.org/doc/install)):

then make your repository (For example myrepo)

go to repository folder (Example myrepo)
and download rhumdapp:
```
cd ~/go/src/github.com/myrepo
git clone https://github.com/marcos9966/rhumdapp.git
cd rhumdapp
```

replace marcos9966 reference with your repository reference contained in the files
main.go and route.go (for example using vi editor)
```
vi main.go
vi route.go
```


download go-bindata utility (used in this project)
```
go get -u github.com/jteeuwen/go-bindata/...
```

download dependecies
```
go get -u github.com/gin-gonic/gin
go get github.com/zserge/webview
go get github.com/phayes/freeport
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

maybe you may see some notice about deprecated calls, however if it terminated without errors, you can run application with
```
cmd/rhumdapp
```
