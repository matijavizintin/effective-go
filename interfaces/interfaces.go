package interfaces
import (
	"net/http"
	"fmt"
	"os"
)

// A channel that sends a notification on each visit.
// (Probably want the channel to be buffered.)
type Chan chan *http.Request

func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ch <- req
	fmt.Fprint(w, "notification sent")
}

// printing channel values
func Exit(ch Chan) {
	for {
		val := <-ch
		fmt.Println(*val)
	}
}


// args server
// note that this is not the classic ServerHTTP function but is using the type HandlerFunc func(ResponseWriter, *Request)
// HandlerFunc has implemented a method func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) that calls the f
// that is this function
// in case of a http server, the function just needs to have this two parameters and its fine
func ArgServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, os.Args)
}