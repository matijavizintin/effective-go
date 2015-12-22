package embedding
import (
	"testing"
	"bufio"
)

func TestEmbedding(t *testing.T) {
	p := make([]byte, 10)

	// note that declaration and initialization is split into two lines for explicitness / clarification

	// init reader and writer
	var reader *bufio.Reader
	reader = bufio.NewReader(nil)

	var writer *bufio.Writer
	writer = bufio.NewWriter(nil)

	// this is just a normal strict with 2 fields that can access only the Read method that is defined
	var rw *ReadWriter1
	rw = &ReadWriter1{reader, writer}
	rw.Read(p)
	//rw.Write(p)		this doesn't work

	// this is the embedding struct that has access to all methods from Reader and Writer
	var rwx *ReadWriterX
	rwx = &ReadWriterX{reader, writer}
	rwx.Read(p)			// note that this method is called on *Reader
	rwx.Write(p)		// note that this method is called on *Writer
}