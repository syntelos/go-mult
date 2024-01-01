/*
 * MULTI/FORMAT I/O
 * Copyright 2023 John Douglas Pritchard, Syntelos
 *
 *
 * References
 * 
 * https://github.com/multiformats/multiformat/blob/master/table.csv
 * https://multiformats.io/
 */
package format

import (
	"fmt"
	"testing"
)

func TestObject(t *testing.T){

	var object Object
	object = object.Define(HeaderBase58)
	fmt.Println(object.String())
}
