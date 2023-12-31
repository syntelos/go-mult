/*
 * MULTI/CODEC I/O
 * Copyright 2023 John Douglas Pritchard, Syntelos
 *
 *
 * References
 * 
 * https://github.com/multiformats/multicodec/blob/master/table.csv
 * https://multiformats.io/
 */
package codec

import (
	"crypto/elliptic"
	"crypto/rand"
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
/*
 */
func TestObject(t *testing.T){
	/*
	 * Use an ECCS as a qualified random number.
	 */
	var curve elliptic.Curve = elliptic.P256()
	_, x, y, er := elliptic.GenerateKey(curve, rand.Reader)
	if nil != er {
		t.Fatalf("test base64 transcoding: error from 'elliptic.GenerateKey': %v", er)
	} else {
		/*
		 * An ECCS public key is a qualified random number.
		 */
		var test_vector_pk_src []byte = elliptic.Marshal(curve, x, y)
		if nil != test_vector_pk_src {
			var test_vector_pk_src_len int = len(test_vector_pk_src)

			var object Object
			object = object.Define(HeaderKeyP256PUB)
			object = object.Encode(test_vector_pk_src)

			var test_vector_pk_dec []byte = object.Decode()
			var test_vector_pk_dec_len int = len(test_vector_pk_dec)
			if test_vector_pk_dec_len != test_vector_pk_src_len {
				t.Errorf("Encoding length found (%d) expected (%d)",test_vector_pk_dec_len,test_vector_pk_src_len)
			}
		} else {
			t.Errorf("Error from 'elliptic.Marshal': %v", er)
		}
	}

}
