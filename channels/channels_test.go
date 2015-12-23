package channels
import (
	"testing"
	"time"
	"fmt"
)

func TestUnbuffered(t *testing.T) {
	ch := make(chan int)		// unbuffered channel - synchronous communication

	go func() {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("go: Sending message")
		ch <- 1
		fmt.Println("go: Reciever recieved the message. Sending complete.")
	}()

	// wait for the goroutine to complete
	fmt.Println("main: Waiting")
	time.Sleep(3000 * time.Millisecond)
	<-ch
	fmt.Println("main: Completed")
}

func TestBuffered(t *testing.T) {
	ch := make(chan int, 10)		// buffered channel - asynchronous communication until the channel is full

	go func() {
		fmt.Println("go: Sending first message")
		ch <- 1
		fmt.Println("go: Sending second message")
		ch <- 2
	}()

	fmt.Println("main: Waiting")
	time.Sleep(1000 * time.Millisecond)
	<-ch
	fmt.Println("main: First message recieved.")
	<-ch
	fmt.Println("main: Second message recieved.")
	fmt.Println("main: Completed")

}
