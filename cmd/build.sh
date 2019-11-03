<<<<<<< HEAD
go-bindata  -prefix res -pkg resmgr -nocompress -o ./internal/pkg/resmgr/assets.go res/...
export GO111MODULE=off
=======
go-bindata -prefix res -pkg resmgr -nocompress -o ./internal/pkg/resmgr/assets.go res/...
>>>>>>> reordering package
go1.13.3 build -o cmd/rhumdapp
