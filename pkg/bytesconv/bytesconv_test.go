package bytesconv

import (
	"fmt"
	"testing"
)

func TestStringToBytes(t *testing.T) {
	str := "hello"
	b := []byte(str)
	toString := BytesToString(b)
	b[0] = '1'
	fmt.Println(toString)
	fmt.Println(string(b))
}
