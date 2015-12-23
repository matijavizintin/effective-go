package recover
import (
	"fmt"
)

type Work struct {
	// dummy
}

func Server(workChan <-chan *Work) {
	for work := range workChan {
		go safelyDo(work)		// execute a goroutine with recover function
	}
}

func safelyDo(work *Work) {
	// a deferred function that can recover a panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic recovered!\n")
		}
	}()

	// run a function that panics
	panicButton(work)
}

func panicButton(work *Work) {
	fmt.Println("Panic button pressed!")
	panic("Panic!")
}
