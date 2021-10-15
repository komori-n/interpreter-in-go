package object

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringHashKey(t *testing.T) {
	a := assert.New(t)

	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnney"}
	diff2 := &String{Value: "My name is johnney"}

	a.Equal(hello1.HashKey(), hello2.HashKey())
	a.Equal(diff1.HashKey(), diff2.HashKey())
	a.NotEqual(hello1.HashKey(), diff1.HashKey())
}
