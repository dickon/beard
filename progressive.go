package beard
import "hash"
import "github.com/bakergo/rollsum"

type Scanner struct {
        window uint
        scanned uint
	hashes [] uint32
	rolling hash.Hash32
}

func New(window uint) Scanner {
	return Scanner{window,0,make([]uint32, 0, 1), rollsum.New(uint32(window))}
}

func (p *Scanner) Scan(data []byte) {
	for i:=0; i<len(data); i++ {
		p.rolling.Write(data[i:i+1])
		p.scanned ++
		if (p.scanned >= p.window)  {
			p.hashes = append(p.hashes, p.rolling.Sum32())
		}
	}
}
