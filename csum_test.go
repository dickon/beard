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


func ExampleProgressive() {
	var scanner = Scanner{2,0,make([]uint32, 3), rollsum.New(uint32(2))}
	scanner.Scan([]byte("AAAA"))
	fmt.Printf("scanned %d hashes %x", scanner.scanned, scanner.hashes)
	// Output: scanned 4 hashes [c50083 c50083 c50083]
}