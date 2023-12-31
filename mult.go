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
	"io"
)
/*
 */
type Object []byte
/*
 * Binary code user interface.
 */
type IO interface {
	/*
	 * The MULT producer is replicating.
	 */
	Write(io.Writer) (error)
	/*
	 * The MULT consumer is validating.
	 */
	Read(io.Reader) (Object, error)
}
/*
 */
func Copy(dst []byte, dx, dz int, src []byte, sx, sz int) ([]byte) {
	for dx < dz && sx < sz {

		dst[dx] = src[sx]
		dx += 1
		sx += 1
	}
	return dst
}
/*
 */
func Concatenate(a, b []byte) ([]byte) {
	var a_len int = len(a)
	if 0 == a_len {
		return b
	} else {
		var b_len int = len(b)
		if 0 == b_len {
			return a
		} else {
			var c_len int = (a_len+b_len)

			var c []byte = make([]byte,c_len)

			c = Copy(c,0,c_len,a,0,a_len)

			c = Copy(c,a_len,c_len,b,0,b_len)

			return c
		}
	}
}
