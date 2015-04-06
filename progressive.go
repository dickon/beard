package beard
import "hash"
import "github.com/bakergo/rollsum"

type Scanner struct {
        window uint
        scanned uint
	hashes [] uint32
	rolling hash.Hash32
	blocks map[uint32] [][]byte
}

func New(window uint) Scanner {
	return Scanner{window,0,make([]uint32, 0, 1), rollsum.New(uint32(window)),
		make(map[uint32] [][]byte)}
}

func (p *Scanner) Scan(data []byte) {
	for i:=range data {
		p.rolling.Write(data[i:i+1])
		p.scanned ++
		if (p.scanned >= p.window)  {
			csum := p.rolling.Sum32()
			p.hashes = append(p.hashes, csum)
			start := i+1 - int(p.window)
			if start >= 0 {
				p.blocks[csum] = append(p.blocks[csum], data[start:i+1])
			}
			
		}
	}
}
