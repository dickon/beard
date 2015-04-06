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
	if _, found := scanner.blocks[0xc50083]; found {
		fmt.Print("found ")
	}
	scanner.Scan([]byte("AA"))
	fmt.Printf("scanned %d hashes %x blocks %d AA blocks %d", scanner.scanned, scanner.hashes, len(scanner.blocks), len(scanner.blocks[0xc50083]))
	// Output: found scanned 4 hashes [c50083 c50083 c50083] blocks 1 AA blocks 1
}

func ExampleProgressive2() {
	scanner := New(2)
	scanner.Scan([]byte("AABAA"))
	fmt.Printf("%x", scanner.hashes)
	// Output: [c50083 c60084 c70084 c50083]
}
