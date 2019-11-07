// This package manage resource data
package resmgr

import (
	"fmt"
)
// structure containing variables to handle ui resource
type Resource struct {
	Root string
}

// get resource file
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
