package reflection

import (
	. "fmt"
	"io"
	"os"
	"bufio"
	"bytes"
	"reflect"
)

type MyInt int

var i int
var j MyInt
var r io.Reader
var w io.Writer
var iface interface{}

func ReflectionImpl() {
	i = 1
	j = 2

	Print(i)
	Print(j)

	// any type which implements "Read(p []byte) (n int, err error)" method can be said that `It implements io.Reader's interface.`
	r = os.Stdin
	r = bufio.NewReader(r)
	r = new(bytes.Buffer)

	iface = r

	hoge, err := os.OpenFile("reflection/hoge.txt", os.O_RDWR, 0)

	if err != nil {
		Print(err)
		return
	}

	// type Reader interface {
	// 	Read(p []byte) (n int, err error)
	// }
	// type Writer interface {
	// 	Write(p []byte) (n int, err error)
	// }

	r = hoge // r contains, schematically, the (value, type) pair, (tty, *os.File)
	w = r.(io.Writer) // assertion. w contains the pair (tty, *os.File)

	Printf("\n")

	// =====================================================

	var x float64 = 3.4

	Println("type:", reflect.TypeOf(x)) // print type of x
	Println("type:", reflect.TypeOf(reflect.TypeOf(x).String())) // print type of "type of x" that is converted to string
	Println("value:", reflect.ValueOf(x).String()) // print value of x's Value

	v := reflect.ValueOf(x)

	Println("type:", v.Type())
	Println("kind is float64:", v.Kind() == reflect.Float64)
	Println("value:", v.Float())

	var y uint8 = 'x'

	v1 := reflect.ValueOf(y)

	Printf("value of y %v\n", y)
	Printf("value of v1 %v\n", v1)
	Println("type:", v1.Type())
	Println("kind is uint8: ", v1.Kind() == reflect.Uint8)
	y = uint8(v1.Uint())
}
