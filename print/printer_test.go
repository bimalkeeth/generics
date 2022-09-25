package print

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func Test_Print_Data(t *testing.T) {
	buf := bytes.NewBufferString("Hello")
	fmt.Println(Identity[io.Reader](io.Reader(buf)))
}

func Test_Print_Data2(t *testing.T) {
	d := Bunch[float64]{123.34, 345.67, 78.98}
	fmt.Println(Bunching(d))

}

func Test_Print_Data3(t *testing.T) {

	type dd struct {
		name string
	}

	a := dd{"ddd"}
	b := dd{"ddd"}

	fmt.Println(Equal(a, b))
}
