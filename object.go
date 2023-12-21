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
	"bytes"
	"fmt"
	"io"
	varint "github.com/syntelos/go-varint"
)
/*
 * Encoded data set content object.
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
func (this Object) Write(w io.Writer) (e error){
	_, e = w.Write(this)
	return e
}
/*
 */
func (this Object) Read(r io.Reader) (Object, error){
	var b io.ByteReader  = r.(io.ByteReader)
	var header uint64
	var er error
	header, er = varint.Read(b)

	if nil != er {
		return nil, fmt.Errorf("Read: %w",er)
	} else {
		this = this.Define(Header(header))

		var payload []byte

		payload, er = io.ReadAll(r)

		if nil != er {

			return nil, fmt.Errorf("Read: %w",er)
		} else {
			this = this.Concatenate(payload)

			return this, nil
		}
	}
}
/*
 * Describe header.
 */
func (this Object) String() string {

	return this.Header().String()
}
/*
 * Resolve object as header extraction.
 */
func (this Object) Header() (header Header) {
	var source *bytes.Buffer = bytes.NewBuffer(this)

	value, er := varint.Read(source)
	if nil != er {
		return 0
	} else {
		return Header(value)
	}
}
/*
 * Define object as header encapsulation.
 */
func (this Object) Define(header Header) (Object) {
	var target *bytes.Buffer = new(bytes.Buffer)

	ct, er := varint.Write(target,uint64(header))
	if nil != er || 0 == ct {
		this = nil
	} else {
		this = target.Bytes()
	}
	return this
}
/*
 * Resolve object content as payload extraction.
 */
func (this Object) Decode() ([]byte) {

	var headcount uint64 = varint.Count(uint64(this.Header()))

	return this[headcount:]
}
/*
 * Define object content as payload encapsulation.
 */
func (this Object) Encode(payload []byte) (Object) {

	return this.Concatenate(payload)
}
/*
 */
func copier(dst []byte, dx, dz int, src []byte, sx, sz int) ([]byte) {
	for dx < dz && sx < sz {

		dst[dx] = src[sx]
		dx += 1
		sx += 1
	}
	return dst
}
func (this Object) Concatenate(b []byte) (Object) {
	var a []byte = this
	var a_len int = len(a)
	if 0 == a_len {
		this = b
	} else {
		var b_len int = len(b)
		if 0 == b_len {
			this = b
		} else {
			var c_len int = (a_len+b_len)

			var c []byte = make([]byte,c_len)

			c = copier(c,0,c_len,a,0,a_len)

			c = copier(c,a_len,c_len,b,0,b_len)

			this = c
		}
	}
	return this
}
