# rhumdapp

rhumdapp is based on [webview](https://github.com/zserge/webview) go library.
The purpose of rhumdapp is to be the skeleton to write desktop application using go, javascript and html.
Currently the project is tested on linux mint distro, so the following instruction are referred to debian based distro.

### Getting started

Install go (debian dist):
```
$ sudo apt install golang-go
```

set path go/bin
```
vi .bashrc
```
at the end of file add follow line
```
export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin
```
Save, quit terminal and reopen it

download go-bindata utility (used in this project)
```
go get -u github.com/jteeuwen/go-bindata/...
```


download zip, extract content and then with terminal go to rhumdapp-master folder.
enable build shell to execute:

```
chmod +x ./cmd/build.sh
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

If you have problem with GTK3 install (requested by package webview)
```
sudo apt-get install build-essential libgtk-3-dev
```
If building project you get the follow error message: "Package webkit2gtk-4.0 was not found in the pkg-config search path" install the package
```
sudo apt-get install libwebkit2gtk-4.0-dev
```
trying again to build project you may have some notice, however you can run application with
```
cmd/rhumdapp
``` 