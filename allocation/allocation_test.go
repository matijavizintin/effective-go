package allocation

import (
	"testing"
	"fmt"
	"sync"
	"reflect"
)

type Allocation struct {
	x     int
	mutex sync.Mutex // zero value of mutex is an unlocked mutex
}

func (a Allocation) String() string {
	return fmt.Sprintf("%s [x=%d, mutex=%v]", reflect.TypeOf(a), a.x, a.mutex)
}

func TestNew(t *testing.T) {
	// it allocates memory, but it doesn't initialize it, memory is zeroed
	// it returns a pointer to the allocated type; this is named zero value
	var v *Allocation = new(Allocation)

	fmt.Printf("%v", v)
}

func TestInitialize(t *testing.T) {
	v := &Allocation{x: 5}                // init pointer to Allocation with x = 5

	fmt.Println(*v)
}

func TestMake(t *testing.T) {
	// make is used to initialize (not zeroed like new) slices, maps and channels
	// this 3 types must be initialized before use
	s := make([]int, 0)		// produces a slice of type int and size 0
	m := make(map[int]string)	// produces a map with key int and value string
	c := make(chan int)		// produces an unbuffered channel of type int

	fmt.Println(s, m, c)

	// otherwise slices and maps can be initialized like this
	s1 := []int{}
	m1 := map[int]string{}

	fmt.Println(s1, m1)
}
