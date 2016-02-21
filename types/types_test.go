package types

import (
	"testing"
	"fmt"
)

func TestInterfaces(t *testing.T) {
	var value interface{}

	value = StringHolder{"abc"}    // Value provided by caller.

	// type switch for each case in a sense converts the value to the type in the case
	switch str := value.(type) {
	case string:
		fmt.Println(str)
	case Stringer:
		fmt.Println(str.String())
	}
}

func TestTypeExtraction(t *testing.T) {
	var value interface{}

	value = StringHolder{"abc"}        // set value to var value

	// extract Stringer type from the value
	str, ok := value.(Stringer)
	if ok {
		fmt.Println("Extracted value: ", str)
	} else {
		fmt.Println("Not of type string")
	}
}

// interface defining String method
type Stringer interface {
	String() string
}

// implementation of Stringer interface
type StringHolder struct {
	X string
}

func (s StringHolder) String() string {
	return s.X
}


