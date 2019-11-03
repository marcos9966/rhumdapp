# rhumdapp

rhumdapp is based on [webview](https://github.com/zserge/webview) go library.
The purpose of rhumdapp is to be the skeleton to write desktop application using go, javascript and html.

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

download go1.13.3
```
go get golang.org/dl/go1.13.3
go1.13.3 download
```

download go-bindata utility (used in this project)
```
go1.13.3 get -u github.com/jteeuwen/go-bindata/...
```


download zip, extract content and then with terminal go to rhumdapp-master folder.
enable build shell to execute:

```
chmod +x ./cmd/build.sh
```

download dependecies
```
<<<<<<< HEAD
go get -u ...
=======
go get -u go get -u github.com/gin-gonic/gin
>>>>>>> reordering package
```
