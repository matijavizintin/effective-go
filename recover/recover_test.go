package recover
import (
	"testing"
	"time"
)

func TestRecover(t *testing.T) {
	ch := make(chan *Work, 10)

	// run a server and give him some work
	go Server(ch)
	ch <- &Work{}
	ch <- &Work{}
	ch <- &Work{}

	time.Sleep(time.Second)
}