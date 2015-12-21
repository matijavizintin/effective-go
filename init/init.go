package init

// first the imported packages are initialized
import "fmt"

// secondly the val declaration is evaluated
var X = 5

// finally the init function is called
func init() {
	X = 6
}

// a file can have multiple init functions
func init() {
	X = 7
}

func PrintState() {
	fmt.Println(X)
}

