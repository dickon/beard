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
	fmt.Printf("out %x", rolling.Sum32())
	rolling.Write([]byte("AA"))
	fmt.Printf(" out %x", rolling.Sum32())
	// Output: out c50083 out 28e0105
}

type Scanner struct {
        window uint
        scanned uint
	first uint32
}

func (p *Scanner) Scan(data []byte) {
	p.scanned += uint(len(data))
	rolling := rollsum.New(2)
	rolling.Write(data[:2])
	p.first = rolling.Sum32()
}

func ExampleProgressive() {
	var scanner = Scanner{2,0, 0}
	scanner.Scan([]byte("AAAA"))
	fmt.Printf("scanned %d first %x", scanner.scanned, scanner.first)
	// Output: scanned 4 hashes [c50083, c50083, c50083]
}