package interfaces
import (
	"testing"
	"net/http"
	"log"


	// note the blank identifier - package is renamed to blank identifier - otherwise the program won't compile
	// because the package is not used
	// this is used for debug
	_ "net/http/pprof"
)

func TestServerAndChannel(t *testing.T) {
	// channel handler
	c := make(Chan, 10)
	http.Handle("/chan", c)

	// args handler
	http.Handle("/args", http.HandlerFunc(ArgServer))

	// run channel listener
	go Exit(c)

	// run server that writes to channel
	err := http.ListenAndServe("localhost:4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
