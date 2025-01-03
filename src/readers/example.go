// package main

// import (
// 	"fmt"
// 	"io"
// 	"strings"
// )

// func main() {
// 	r := strings.NewReader("Hello, Reader!")

// 	b := make([]byte, 8)
// 	for {
// 		n, err := r.Read(b)
// 		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
// 		fmt.Printf("b[:n] = %q\n", b[:n])
// 		if err == io.EOF {
// 			break
// 		}
// 	}
// }

// package main

// import "golang.org/x/tour/reader"

// type MyReader struct{}

// // TODO: Add a Read([]byte) (int, error) method to MyReader.
// func (x MyReader) Read(b []byte) (int, error) {
// 	for i := range b {
// 		b[i] = 'A'
// 	}
// 	return len(b), nil

// }

// func main() {
// 	reader.Validate(MyReader{})
// }

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	for i := 0; i < n; i++ {
		if (p[i] >= 'A' && p[i] <= 'Z') || (p[i] >= 'a' && p[i] <= 'z') {
			// Determine the base ('A' for uppercase, 'a' for lowercase)
			base := byte('A')
			if p[i] >= 'a' {
				base = 'a'
			}
			// Apply ROT13 transformation
			p[i] = (p[i]-base+13)%26 + base
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
