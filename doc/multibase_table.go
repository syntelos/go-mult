/*
 * MULTIBASE TABLE I/O
 * Copyright 2023 John Douglas Pritchard, Syntelos
 *
 * References
 * 
 * [MULT] https://github.com/syntelos/go-mult
 * [MUCT] https://github.com/multiformats/multicodec/blob/master/table.csv
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)
/*
 */
func usage(){
	fmt.Fprintln(os.Stderr,`
Synopsis

    table [print|enumerate|list] [fwd|rev|ext]

Description

    Read table, and print or enumerate.  Optionally invert
    or extend identity-code order of relationship.

`)
	os.Exit(1)
}
/*
 */
type Line struct {
	unicode string
	character string
	encoding string
	description string
	status string
	cid string
}
func trim(src []byte) (string) {
	var cnt int = len(src)
	if 0 < cnt {
		var last int = (cnt-1)
		if ws(src[last]) {
			var tail int = (last-1)
			if 0 <= tail {
				for ws(src[tail]) {
					last = tail
					tail -= 1;
					if 0 > tail {
						return ""
					}
				}
			}
			return string(src[0:last])
		} else {
			return string(src)
		}
	}
	return string(src)
}
func rewrite(src []byte) (string) {
	var ch byte
	var idx, cnt int = 0, len(src)

	var tgt []byte = make([]byte,cnt)

	for ; idx < cnt; idx++ {
		ch = src[idx]
		switch ch {
		case '"':
			tgt[idx] = '\''
		case ',':
			tgt[idx] = ';'
		default:
			tgt[idx] = ch
		}
	}

	return string(tgt)
}
/*
 * Text span operator for character classes.
 */
type cc func (byte)(bool)
/*
 * Text span counter for character classes.
 */
func class (src []byte, ofs int, len int, clop cc) (spx int) {
	spx = ofs

	for ; ofs < len; ofs++ {

		if clop(src[ofs]) {

			spx = ofs
		} else {
			return spx
		}
	}
	return spx
}
/*
 * CSV content character class.
 */
func id (ch byte) (bool) {
	switch {
	case 'a' <= ch && 'z' >= ch :
		return true
	case 'A' <= ch && 'Z' >= ch :
		return true
	case '0' <= ch && '9' >= ch :
		return true
	case '_' == ch || '-' == ch  || '+' == ch  || '.' == ch :
		return true
	case '(' == ch || ')' == ch  || '[' == ch  || ']' == ch :
		return true
	case ' ' == ch:
		return true
	default:
		return false
	}
}
/*
 * CSV delimeter character class.
 */
func ws (ch byte) (bool) {
	switch {
	case '\r' == ch || '\n' == ch || '\t' == ch || ' ' == ch || ',' == ch:
		return true
	default:
		return false
	}
}
/*
 * CSV line reader.
 */
func (this Line) read(inl []byte) (Line) {

	var begin, end, inz int = 0, 0, len(inl)
	{
		end = class(inl,begin,inz,id)+1
	}
	this.unicode = trim(inl[begin:end])
	{
		begin = class(inl,end,inz,ws)+1

		end = class(inl,begin,inz,id)+1
	}
	this.character = trim(inl[begin:end])
	{
		begin = class(inl,end,inz,ws)+1

		end = class(inl,begin,inz,id)+1
	}
	this.encoding = trim(inl[begin:end])
	{
		begin = class(inl,end,inz,ws)+1

		end = class(inl,begin,inz,id)+1	
	}
	this.description = rewrite(inl[begin:end])
	{
		begin = class(inl,end,inz,ws)+1
	}
	this.status = trim(inl[begin:])

	this.cid = Camel(this.encoding)

	return this
}
func Camel(src_str string) (string) {
	var dst strings.Builder
	var src []byte = []byte(src_str)
	var src_len int = len(src)

	var w bool = true
	var x, z int = 0, src_len
	var y byte

	for ; x < z; x++ {
		y = src[x]
		if w {
			if 'a' <= y && 'z' >= y {
				y -= 'a'
				y += 'A'

				dst.WriteByte(y)

				w = false
			} else {
				dst.WriteByte(y)
			}
		} else if '-' == y || '_' == y {
			w = true
		} else if '0' <= y && '9' >= y {
			w = true
			dst.WriteByte(y)
		} else {
			dst.WriteByte(y)
		}
	}
	return dst.String()
}
func (this Line) print(order Order){

	switch order {
	case Fwd:
		fmt.Printf("%s\t%s\t%s\n",this.encoding,this.character,this.description)

	case Rev:
		fmt.Printf("%s\t%s\t%s\n",this.character,this.encoding,this.description)

	default:
		fmt.Printf("%s\t%s\t%s\t%s\n",this.unicode,this.character,this.encoding,this.description)
	}
}
func (this Line) enumerate(order Order){

	switch order {
	case Fwd:
		fmt.Printf("\tcase Identifier%s:\n\t\treturn uint32(Header%s)\n",this.cid,this.cid)

	default:
		fmt.Printf("\tcase Header%s:\n\t\treturn string(Identifier%s)\n",this.cid,this.cid)
	}
}
func (this Line) list(order Order){
	switch order {
	case Fwd:
		fmt.Printf("\tHeader%s\tHeader\t=\t'%s'\n",this.cid,this.character)
	case Rev:
		fmt.Printf("\tIdentifier%s\tIdentifier\t=\t\"%s\"\n",this.cid,this.cid)

	default:
		fmt.Printf("\tconst\t%s\tuint32\t=\t'%s'\n",this.cid,this.character)
	}
}
/*
 */
type Table struct {
	filename string
	records []Line
}

var table *Table = new(Table)

func (this *Table) size() (z int){

	return len(this.records)
}
func (this *Table) read(filename string) (e error){
	this.filename = filename

	var file *os.File
	file, e = os.Open(filename)
	if nil != e {
		e = fmt.Errorf("Error opening '%s': %w",filename,e)
		return e
	} else {
		defer file.Close()

		var reader *bufio.Reader = bufio.NewReader(file)

		var inl []byte
		var isp, once bool = false, true
		var lin Line
		inl, isp, e = reader.ReadLine()

		for true {
			if nil != e {
				if io.EOF == e {

					return nil
				} else {
					return fmt.Errorf("Error reading '%s': %w",filename,e)
				}
			} else if isp {
				return fmt.Errorf("Error reading '%s'.",filename)
			} else {
				if once {
					once = false
				} else {
					this.records = append(this.records,lin.read(inl))
				}
				inl, isp, e = reader.ReadLine()
			}
		}
		return nil
	}
}
func (this *Table) print(order Order){

	var count int = table.size()
	fmt.Printf("# %s %d\n",this.filename,count)

	var index int = 0
	for ; index < count; index++ {
		this.records[index].print(order)
	}
}
func (this *Table) enumerate(order Order){

	var count int = table.size()
	var index int = 0
	for ; index < count; index++ {
		this.records[index].enumerate(order)
	}
}
func (this *Table) list(order Order){

	var count int = table.size()
	var index int = 0
	for ; index < count; index++ {
		this.records[index].list(order)
	}
}
/*
 */
const location_rel string = "multibase_table.csv"
const location_doc string = "doc/multibase_table.csv"

func (this *Table) location() (string, error) {
	_, er := os.Stat("doc")
	if nil == er {
		_, er := os.Stat(location_doc)
		if nil == er {
			return location_doc, nil
		} else {
			return "", er
		}
	} else {
		_, er := os.Stat(location_rel)
		if nil == er {
			return location_rel, nil
		} else {
			return "", er
		}
	}
}
type Order uint8
const (
	Fwd Order = iota
	Rev
	Ext
)
func (this Order) Parse(order string) (Order) {
	switch order {
	case "fwd":
		this = Fwd
	case "rev":
		this = Rev
	default:
		this = Ext
	}
	return this
}
/*
 */
func main(){
	var argc int = len(os.Args)
	var argx int = 1
	if argx < argc {
		var opr string = os.Args[argx]

		var fin string
		var err error
		fin, err = table.location()
		if nil != err {
			fmt.Fprintf(os.Stderr,"table: file not found '%s'.\n",location_doc);
		} else {
			var order Order = Ext
			argx += 1
			if argx < argc {
				order = order.Parse(os.Args[argx])
			}

			switch opr {

			case "print":
				e := table.read(fin)
				if nil != e {
					fmt.Fprintf(os.Stderr,"table: %v\n",e);
					os.Exit(1)
				} else {
					table.print(order)

					os.Exit(0)
				}

			case "enumerate":
				e := table.read(fin)
				if nil != e {
					fmt.Fprintf(os.Stderr,"table: %v\n",e);
					os.Exit(1)
				} else {
					table.enumerate(order)

					os.Exit(0)
				}

			case "list":
				e := table.read(fin)
				if nil != e {
					fmt.Fprintf(os.Stderr,"table: %v\n",e);
					os.Exit(1)
				} else {
					table.list(order)

					os.Exit(0)
				}

			default:
				usage()
			}
		}
	} else {
		usage()
	}
}
