package closure

import "net/http"

var sem = make(chan int, 10)

func process(req *http.Request) {
	// dummy
}

func Serve(queue chan *http.Request) {
	for req := range queue {
		sem <- 1
		go func() {
			// NOTE: in this case req variable will be shared across all request due to closures - this is a bug
			// we don't want that
			process(req)
			<-sem
		}()
	}
}

func Serve1(queue chan *http.Request) {
	for req := range queue {
		sem <- 1
		go func(req *http.Request) {
			// here the req var is passed as an argument so the req in a for loop is not included in the closure
			process(req)
			<-sem
		}(req)
	}
}

func Serve2(queue chan *http.Request) {
	for req := range queue {
		req := req // Create new instance of req for the goroutine.
		sem <- 1
		go func() {
			// since this is a copy of the original req variable the instance in the loop won't be included in the closure
			process(req)
			<-sem
		}()
	}
}
