package beard

import "hash"
import "bytes"
import "github.com/bakergo/rollsum"
import "hash/adler32"
import "errors"

// Scanner scans data and produce hashes
type Scanner struct {
	window  uint
	scanned uint
	hashes  []uint32
	rolling hash.Hash32
	blocks  map[uint32][][]byte
	blockindex []*[]byte
}

func NewScanner(window uint) Scanner {
	return Scanner{window, 0, make([]uint32, 0, 1), rollsum.New(uint32(window)),
		make(map[uint32][][]byte), make([]*[]byte, 0, 1)}
}

func (p *Scanner) Scan(data []byte) {
	for i := range data {
		p.rolling.Write(data[i : i+1])
		p.scanned++
		if p.scanned >= p.window {
			csum := p.rolling.Sum32()
			p.hashes = append(p.hashes, csum)
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
		if bytes.Compare(value, block) == 0 {
			return
		}
	}
	p.blocks[csum] = append(p.blocks[csum], block)
	p.blockindex = append(p.blockindex, &p.blocks[csum][len(p.blocks[csum])-1])
}

type Content struct {
	block uint
	location uint
}

func (p *Scanner) Encode(data []byte) ([] Content, error) {
	csum := adler32.Checksum(data[0:p.window])
	found := false
	index := uint(0)
	for _, candidate := range p.blocks[csum] {
		if bytes.Compare(candidate, data) == 0 {
			found = true
			break
		}
	}
	encoding := make([]Content, 1, 1)
	if !found {
		return encoding, errors.New("block not found")
	}
	encoding[0].block = index
	encoding[0].location = 0
	return encoding, nil
}