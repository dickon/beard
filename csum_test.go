package beard

import "github.com/bakergo/rollsum"
import "fmt"

func ExampleSmall() {
	var rolling = rollsum.New(16)
	rolling.Write([]byte("AAAA"))
	fmt.Printf("out %x", rolling.Sum32())
	// Output: out 28e0105
}

func ExampleSmallInc() {
	var rolling = rollsum.New(16)
	rolling.Write([]byte("AA"))
	rolling.Write([]byte("AA"))
	fmt.Printf("out %x", rolling.Sum32())
	// Output: out 28e0105
}

type Scanner struct {
        window uint
        scanned uint
}

func (p *Scanner) Scan(data []byte) {
}

func ExampleProgressive() {
	var scanner = Scanner{2,0}
	scanner.Scan([]byte("AAAA"))
	fmt.Printf("scanned %d", scanner.scanned)
	// Output: out 4
}