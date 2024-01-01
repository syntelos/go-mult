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
	"io"
	mult "github.com/syntelos/go-mult"
	"unicode/utf8"
)
/*
 * Multiformat header as UNICODE rune.
 */
type Header uint32
/*
 * Multiformat identifier as GOPL ID.
 */
type Identifier string
/*
 * Multiformat constants include GOPL domain UNKNOWNs.
 */
const (
	HeaderUNKNOWN		Header = 0
	HeaderBase2	Header	=	'0'
	HeaderNone	Header	=	'1'
	HeaderBase8	Header	=	'7'
	HeaderBase10	Header	=	'9'
	HeaderBase16	Header	=	'f'
	HeaderBase16Upper	Header	=	'F'
	HeaderBase32Hex	Header	=	'v'
	HeaderBase32Hexupper	Header	=	'V'
	HeaderBase32Hexpad	Header	=	't'
	HeaderBase32Hexpadupper	Header	=	'T'
	HeaderBase32	Header	=	'b'
	HeaderBase32Upper	Header	=	'B'
	HeaderBase32Pad	Header	=	'c'
	HeaderBase32Padupper	Header	=	'C'
	HeaderBase32Z	Header	=	'h'
	HeaderBase36	Header	=	'k'
	HeaderBase36Upper	Header	=	'K'
	HeaderBase58	Header	=	'z'
	HeaderBase58Flickr	Header	=	'Z'
	HeaderBase64	Header	=	'm'
	HeaderBase64Pad	Header	=	'M'
	HeaderBase64Url	Header	=	'u'
	HeaderBase64Urlpad	Header	=	'U'
	HeaderProquint	Header	=	'p'

	IdentifierUNKNOWN			Identifier	= ""
	IdentifierBase2	Identifier	=	"Base2"
	IdentifierNone	Identifier	=	"None"
	IdentifierBase8	Identifier	=	"Base8"
	IdentifierBase10	Identifier	=	"Base10"
	IdentifierBase16	Identifier	=	"Base16"
	IdentifierBase16Upper	Identifier	=	"Base16Upper"
	IdentifierBase32Hex	Identifier	=	"Base32Hex"
	IdentifierBase32Hexupper	Identifier	=	"Base32Hexupper"
	IdentifierBase32Hexpad	Identifier	=	"Base32Hexpad"
	IdentifierBase32Hexpadupper	Identifier	=	"Base32Hexpadupper"
	IdentifierBase32	Identifier	=	"Base32"
	IdentifierBase32Upper	Identifier	=	"Base32Upper"
	IdentifierBase32Pad	Identifier	=	"Base32Pad"
	IdentifierBase32Padupper	Identifier	=	"Base32Padupper"
	IdentifierBase32Z	Identifier	=	"Base32Z"
	IdentifierBase36	Identifier	=	"Base36"
	IdentifierBase36Upper	Identifier	=	"Base36Upper"
	IdentifierBase58	Identifier	=	"Base58"
	IdentifierBase58Flickr	Identifier	=	"Base58Flickr"
	IdentifierBase64	Identifier	=	"Base64"
	IdentifierBase64Pad	Identifier	=	"Base64Pad"
	IdentifierBase64Url	Identifier	=	"Base64Url"
	IdentifierBase64Urlpad	Identifier	=	"Base64Urlpad"
	IdentifierProquint	Identifier	=	"Proquint"

)
/*
 * Identify code.
 */
func (code Header) String() (string) {
	switch code {
	case HeaderBase2:
		return string(IdentifierBase2)
	case HeaderNone:
		return string(IdentifierNone)
	case HeaderBase8:
		return string(IdentifierBase8)
	case HeaderBase10:
		return string(IdentifierBase10)
	case HeaderBase16:
		return string(IdentifierBase16)
	case HeaderBase16Upper:
		return string(IdentifierBase16Upper)
	case HeaderBase32Hex:
		return string(IdentifierBase32Hex)
	case HeaderBase32Hexupper:
		return string(IdentifierBase32Hexupper)
	case HeaderBase32Hexpad:
		return string(IdentifierBase32Hexpad)
	case HeaderBase32Hexpadupper:
		return string(IdentifierBase32Hexpadupper)
	case HeaderBase32:
		return string(IdentifierBase32)
	case HeaderBase32Upper:
		return string(IdentifierBase32Upper)
	case HeaderBase32Pad:
		return string(IdentifierBase32Pad)
	case HeaderBase32Padupper:
		return string(IdentifierBase32Padupper)
	case HeaderBase32Z:
		return string(IdentifierBase32Z)
	case HeaderBase36:
		return string(IdentifierBase36)
	case HeaderBase36Upper:
		return string(IdentifierBase36Upper)
	case HeaderBase58:
		return string(IdentifierBase58)
	case HeaderBase58Flickr:
		return string(IdentifierBase58Flickr)
	case HeaderBase64:
		return string(IdentifierBase64)
	case HeaderBase64Pad:
		return string(IdentifierBase64Pad)
	case HeaderBase64Url:
		return string(IdentifierBase64Url)
	case HeaderBase64Urlpad:
		return string(IdentifierBase64Urlpad)
	case HeaderProquint:
		return string(IdentifierProquint)
	default:
		return string(IdentifierUNKNOWN)
	}
}
/*
 * Encode identifier.
 */
func (id Identifier) Code() (uint32) {
	switch id {
	case IdentifierBase2:
		return uint32(HeaderBase2)
	case IdentifierNone:
		return uint32(HeaderNone)
	case IdentifierBase8:
		return uint32(HeaderBase8)
	case IdentifierBase10:
		return uint32(HeaderBase10)
	case IdentifierBase16:
		return uint32(HeaderBase16)
	case IdentifierBase16Upper:
		return uint32(HeaderBase16Upper)
	case IdentifierBase32Hex:
		return uint32(HeaderBase32Hex)
	case IdentifierBase32Hexupper:
		return uint32(HeaderBase32Hexupper)
	case IdentifierBase32Hexpad:
		return uint32(HeaderBase32Hexpad)
	case IdentifierBase32Hexpadupper:
		return uint32(HeaderBase32Hexpadupper)
	case IdentifierBase32:
		return uint32(HeaderBase32)
	case IdentifierBase32Upper:
		return uint32(HeaderBase32Upper)
	case IdentifierBase32Pad:
		return uint32(HeaderBase32Pad)
	case IdentifierBase32Padupper:
		return uint32(HeaderBase32Padupper)
	case IdentifierBase32Z:
		return uint32(HeaderBase32Z)
	case IdentifierBase36:
		return uint32(HeaderBase36)
	case IdentifierBase36Upper:
		return uint32(HeaderBase36Upper)
	case IdentifierBase58:
		return uint32(HeaderBase58)
	case IdentifierBase58Flickr:
		return uint32(HeaderBase58Flickr)
	case IdentifierBase64:
		return uint32(HeaderBase64)
	case IdentifierBase64Pad:
		return uint32(HeaderBase64Pad)
	case IdentifierBase64Url:
		return uint32(HeaderBase64Url)
	case IdentifierBase64Urlpad:
		return uint32(HeaderBase64Urlpad)
	case IdentifierProquint:
		return uint32(HeaderProquint)
	default:
		return uint32(HeaderUNKNOWN)
	}
}
/*
 * Encoded data set content object.
 */
type Object []byte
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
	var header byte
	var er error
	header, er = b.ReadByte()

	if nil != er {
		return nil, fmt.Errorf("Read: %w",er)
	} else {
		this = this.Define(Header(uint32(header)))

		var payload []byte

		payload, er = io.ReadAll(r)

		if nil != er {

			return nil, fmt.Errorf("Read: %w",er)
		} else {
			this = mult.Concatenate(this,payload)

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
	var r rune
	var z int
	r,z = utf8.DecodeRune(this)
	if 0 < z {
		return Header(r)
	} else {
		return HeaderUNKNOWN
	}
}
/*
 * Define object as header encapsulation.
 */
func (this Object) Define(header Header) (Object) {
	if 0 != len(this) {
		this = make([]byte,0)
	}
	this = utf8.AppendRune(this,rune(header))
	return this
}
/*
 * Resolve object content as payload extraction.
 */
func (this Object) Decode() ([]byte) {
	var z int
	_,z = utf8.DecodeRune(this)
	if 0 < z {
		return this[z:]
	} else {
		return nil
	}
}
/*
 * Define object content as payload encapsulation.
 */
func (this Object) Encode(payload []byte) (Object) {

	return mult.Concatenate(this,payload)
}
