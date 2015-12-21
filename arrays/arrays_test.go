package arrays
import (
	"testing"
	"fmt"
)

func TestArrays(t *testing.T) {
	array := [3]float64{7.0, 8.5, 9.1}		// this is an array (not a slice) of 3 elements
	x := Sum(&array)                        // Note the explicit address-of operator
	fmt.Println(x)
}

// note that the array in passed as a pointer - otherwise go makes a copy of the passed parameter - the whole array
func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}