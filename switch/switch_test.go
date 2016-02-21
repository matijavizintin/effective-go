package switchs

import (
	"testing"
	"fmt"
	"strconv"
)

func TestSwitch1(t *testing.T) {
	v := 10

	res := 0
	switch {        // switch with no parameter to switch on
	case 0 <= v && v < 10:
		res = 1
	case 10 <= v && v < 20:
		res = 2
	default:
		res = 3
	}

	fmt.Println(res)
}

func TestSwitch2(t *testing.T) {
	v := 5

	res := 0
	switch v {
	case 0, 1, 2, 3, 4:        // list of case parameters
		res = 1
	case 5, 6, 7, 8:
		res = 2
	}

	fmt.Println(res)
}

func TestBreak(t *testing.T) {
	Loop:
	for i := 0; i < 10; i++ {
		switch {
		case i == 2:
			if true {
				break                // break the switch early
			}
			// some more code to be executed
		case i == 3:
			break Loop                // break to Loop label
		}
	}
}

func TestTypeSwitch(t *testing.T) {
	var v interface{}        // type switch works on interface type
	v = 10

	switch v.(type) {                // switch on v type
	case string:
		fmt.Println(v)
	case int:
		fmt.Println(strconv.FormatInt(v.(int64), 10))
	}
}
