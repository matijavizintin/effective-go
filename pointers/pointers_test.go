package pointers

import (
	"fmt"
	"testing"
)

type MyType int

func TestPointers(t *testing.T) {
	var m = MyType(5)
	fmt.Println(m)

	s := m.String() // this is actually rewritten to (&m).String()
	fmt.Println(s)
}

// function defined on a pointer to MyType
// value methods can be invoked on values and pointers
// pointer methods can only be invoked on pointers
// but when a value is addressable go adds the address operator
func (m *MyType) String() string {
	return fmt.Sprintln(*m)
}
