// This package manage resource data
package resmgr

import (
	"fmt"
)

type Resource struct {
	Root string
}

func (rs Resource) Get(url string) ([]byte, error) {
	fname := rs.Root+url
	fmt.Println(fname)
	data, err := Asset(fname)
	if err != nil {
		return nil,err
	}
	fmt.Println(fname)
	return data,nil
}
