package beard


import "github.com/bakergo/rollsum"
import "testing"
import "fmt"

func TestSmall(t *testing.T) {
	var rolling = rollsum.New(16)
	_, err := rolling.Write([]byte("AAAA"))
	if err != nil {
		t.Error(err)
	}
	if rolling.Sum32() != 0x28e0105 {
		t.Error("mismatch")
	}
}

func TestSmallInc(t *testing.T) {
	var rolling = rollsum.New(16)
	_, err := rolling.Write([]byte("AA"))
	if err != nil {
		t.Error(err)
	}
	if rolling.Sum32() != 12910723 {
		t.Error("bad csum")
	}
	_, err2 :=rolling.Write([]byte("AA"))
	if err2 != nil {
		t.Error(err2)
	}
	if rolling.Sum32() != 42860805 {
		t.Error("bad csum")
	}
}

func TestProgressive(t *testing.T) {
	scanner := NewScanner(2)
	scanner.Scan([]byte("AA"))
	_, found := scanner.blocks[12910723]
	if !found {
		t.Error("did not find expected block")
	}
	if len(scanner.blockindex) !=1 || string((*scanner.blockindex[0]).data) != "AA" {
		t.Error("could not retrieve block")
	}
	encoding, err := scanner.Encode([]byte("AA"))
	if err != nil {
		t.Error("conversion failed")
	}
	if fmt.Sprint(encoding) != "[{0 0}]" {
		t.Error("unexpected encoding");
	}
	scanner.Scan([]byte("AA"))
	if len(scanner.blockindex) != 1 {
		t.Error("unexpected index size")
	}
	if scanner.scanned != 4 {
		t.Error("scanned wrong")
	}
	if scanner.hashes[0] != 12910723 || scanner.hashes[1] != 12910723 || scanner.hashes[2] != 12910723 {
		t.Error("hashes unexpected")
	}
	if len(scanner.blocks[12910723]) != 1 {
		t.Error("miscount")
	}
}

func TestProgressive2(t *testing.T) {
	scanner := NewScanner(2)
	scanner.Scan([]byte("AABAA"))
	if scanner.hashes[0] != 12910723 || scanner.hashes[1] != 12976260 || scanner.hashes[2] != 13041796 || scanner.hashes[3] != 12910723 {
		t.Error("hashes unexpected")
	}
	if len(scanner.blocks[12910723]) != 1 {
		t.Error("miscount")
	}
	if len(scanner.blocks[12976260]) != 1 {
		t.Error("miscount")
	}
	if len(scanner.blockindex) < 2 || string((*scanner.blockindex[1]).data) != "AB" {
		t.Error("could not retrieve block")
	}
	if len(scanner.blockindex) < 3 || string((*scanner.blockindex[2]).data) != "BA" {
		t.Error("could not retrieve block")
	}
	encoding, err := scanner.Encode([]byte("AABAAA"))
	if err != nil {
		t.Error("conversion failed")
	}
	if fmt.Sprint(encoding) != "[{0 0} {2 2} {0 4}]" {
		t.Error("unexpected encoding");
	}
}
