package beard
import "hash"

type Scanner struct {
        window uint
        scanned uint
	hashes [] uint32
	rolling hash.Hash32
}

func (p *Scanner) Scan(data []byte) {
	for i:=0; i<len(data); i++ {
		p.rolling.Write(data[i:i+1])
		p.scanned ++
		if (p.scanned >= p.window)  {
			p.hashes[p.scanned-p.window] = p.rolling.Sum32()
		}
	}
}
