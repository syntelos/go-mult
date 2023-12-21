MULTI/FORMAT/CODEC I/O

  /*
   * Encoded data set content object.
   */
  type Object []byte
  /*
   * Binary code user interface.
   */
  type IO interface {

	  Write(io.Writer) (error)

	  Read(io.Reader) (error)
  }


References

  [MULF] https://multiformats.io/
  [MUCT] https://github.com/multiformats/multicodec/blob/master/table.csv

