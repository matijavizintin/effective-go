package defer1
import (
	"testing"
	"fmt"
)

func TestDefer(t *testing.T) {
	for i := 0; i < 10; i++ {
		// defer pushes the method execution on the (LIFO) stack and the calls are executed just before the function
		// returns; executing this test will print the values in reverse order
		defer fmt.Printf("%d ", i)
	}

	fmt.Println("This is printed first!")
}
