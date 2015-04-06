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
	scanner := New(2)
	scanner.Scan([]byte("AA"))
	_, found := scanner.blocks[0xc50083]
	scanner.Scan([]byte("AA"))
	fmt.Print(found, scanner.scanned, scanner.hashes, len(scanner.blocks), len(scanner.blocks[12910723]))
	// Output: true 4 [12910723 12910723 12910723] 1 1
}

func ExampleProgressive2() {
	scanner := New(2)
	scanner.Scan([]byte("AABAA"))
	fmt.Printf("%x", scanner.hashes)
	// Output: [c50083 c60084 c70084 c50083]
}
