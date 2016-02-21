package variadic

import (
	"fmt"
	"testing"
)

func TestVariadic(t *testing.T) {
	min := Min(10, 4, 5, 7, 8)
	fmt.Println(min)
}

func TestVariadic2(t *testing.T) {
	Println(1, "asdf", float64(-7.5), complex128(-4.3123i), nil, SomeStruct{1, "x"})
}

// this function can take a variable number of arguments; a basically acts like variable of type []int
func Min(a ...int) int {
	min := int(^uint(0) >> 1) // largest int
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

// this function can take also a variable number of arguments but in contrast of the upper function they can be of
// any type. go doesn't have generics (or this is at least what i understood) so interface is like any type
func Println(v ...interface{}) {
	// note the v... syntax - this means that v is passed as a list of arguments and not as a single argument - a slice
	fmt.Println(2, fmt.Sprintln(v...)) // Output takes parameters (int, string)
}

type SomeStruct struct {
	p1 int
	p2 string
}

// custom toString method for SomeStruct
func (s SomeStruct) String() string {
	return fmt.Sprintf("p1: %T %+v; p2: %T %+v", s.p1, s.p1, s.p2, s.p2)
}
