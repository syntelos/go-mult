/*
 * MULTI/FORMAT/CODEC I/O
 * Copyright 2023 John Douglas Pritchard, Syntelos
 *
 *
 * References
 * 
 * https://github.com/multiformats/multicodec/blob/master/table.csv
 * https://multiformats.io/
 */
package mult

import (
	"fmt"
	"testing"
)
/*
 */
func TestTable(t *testing.T){
	var cc uint64 = IdentifierKeyP256PUB.Code()
	var id string = HeaderKeyP256PUB.String()

	fmt.Printf("%s\t0x%08X\n",id,cc)
}
