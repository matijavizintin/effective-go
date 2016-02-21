package embedding

import "bufio"

// two interfaces
type Reader interface {
	Read(p []byte) (n int, err error)
}
type Writer interface {
	Write(p []byte) (n int, err error)
}

// embedding interface Reader and Writer
// ReadWriter is the interface that combines the Reader and Writer interfaces.
type ReadWriter interface {
	Reader
	Writer
}

// embedding struct that embeds two structs
type ReadWriterX struct {
	*bufio.Reader
	*bufio.Writer
}

// the '1' at the end is used to not conflict with the interface
// this here is just to show the lower forwarding method
type ReadWriter1 struct {
	reader *bufio.Reader
	writer *bufio.Writer
}

// the Read method on the bufio.Reader is used through the ReadWriter1 struct that embeds the reader
// NOTE that when the Read method of the bufio.ReadWriter is invoked it has the same effect as this method, the actual
// receiver is the reader field so in basically invoked on the Reader struct
func (rw *ReadWriter1) Read(p []byte) (n int, err error) {
	return rw.reader.Read(p)
}

type dummyReader struct {
}

func (d dummyReader) Read(p []byte) (n int, err error) {
	p[0] = 'a'
	return 1, nil
}

type dummyWriter struct {
}

func (d dummyWriter) Write(p []byte) (n int, err error) {
	p[0] = 'a'
	return 1, nil
}
