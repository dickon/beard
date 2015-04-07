package beard

import "hash"
import "bytes"
import "github.com/bakergo/rollsum"
import "hash/adler32"
import "errors"

// BlockRecord stores a block and gives its index
type BlockRecord struct {
	data [] byte
	index uint
}

// Scanner scans data and produce hashes
type Scanner struct {
	window  uint
	scanned uint
	rolling hash.Hash32
	blocks  map[uint32][]BlockRecord
	blockindex []*BlockRecord
}

func NewScanner(window uint) Scanner {
	return Scanner{window, 0, rollsum.New(uint32(window)),
		make(map[uint32][]BlockRecord), make([]*BlockRecord, 0, 1)}
}

func (p *Scanner) Scan(data []byte) {
	for i := range data {
		p.rolling.Write(data[i : i+1])
		p.scanned++
		if p.scanned >= p.window {
			csum := p.rolling.Sum32()
			start := i + 1 - int(p.window)
			if start >= 0 {
				p.Store(csum, data[start:i+1])
			}

		}
	}
}

func (p *Scanner) Store(csum uint32, block []byte) {
	blocklist := p.blocks[csum]
	for _, value := range blocklist {
		if bytes.Compare(value.data, block) == 0 {
			return
		}
	}
	index := uint(len(p.blockindex))
	p.blocks[csum] = append(p.blocks[csum], BlockRecord{block, index})
	p.blockindex = append(p.blockindex, &p.blocks[csum][len(p.blocks[csum])-1])
}

type Content struct {
	block uint
	location uint
}

func (p *Scanner) Encode(data []byte) ([] Content, error) {
	encoding := make([]Content, 0, 1)
	for i :=0; i<len(data); i+= int(p.window) {
		dataslice := data[i:i+int(p.window)]
		csum := adler32.Checksum(dataslice)
		found := false
		index := uint(0)
		for _, candidate := range p.blocks[csum] {
			if bytes.Compare(candidate.data, dataslice) == 0 {
				found = true
				index = candidate.index
				break
			}
		}
		if !found {
			return encoding, errors.New("block not found")
		}

		encoding = append( encoding, Content{index, uint(i)})
	}
	return encoding, nil
}