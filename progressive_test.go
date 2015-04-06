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
	first := rolling.Sum32()
	rolling.Write([]byte("AA"))
	fmt.Print(first, rolling.Sum32())
	// Output: 12910723 42860805
}

func ExampleProgressive() {
	scanner := New(2)
	scanner.Scan([]byte("AA"))
	_, found := scanner.blocks[12910723]
	scanner.Scan([]byte("AA"))
	fmt.Print(found, scanner.scanned, scanner.hashes, len(scanner.blocks), len(scanner.blocks[12910723]))
	// Output: true 4 [12910723 12910723 12910723] 1 1
}

func ExampleProgressive2() {
	scanner := New(2)
	scanner.Scan([]byte("AABAA"))
	fmt.Print(scanner.hashes)
	// Output: [12910723 12976260 13041796 12910723]
}
