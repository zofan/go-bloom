package bloom

import (
	"github.com/zofan/go-bitset"
)

type Bloom struct {
	bitset *bitset.BitSet

	keys int
}

func New(bs *bitset.BitSet, keys int) *Bloom {
	return &Bloom{
		bitset: bs,
		keys:   keys,
	}
}

func (b *Bloom) Test(data []byte) bool {
	for n := 0; n < b.keys; n++ {
		if !b.bitset.Test(hash(data, n) % b.bitset.Size()) {
			return false
		}
	}

	return true
}

func (b *Bloom) Add(data []byte) error {
	for n := 0; n < b.keys; n++ {
		err := b.bitset.Set(hash(data, n) % b.bitset.Size())
		if err !=  nil {
			return err
		}
	}

	return nil
}

func hash(data []byte, i int) uint64 {
	data = append(data, byte(i))

	var hash uint64 = 14695981039346656037

	for _, b := range data {
		hash ^= uint64(b)
		hash += (hash << 8) + (hash << 16) + (hash << 24) + (hash << 32) + (hash << 40) + (hash << 48) + (hash << 56)
	}

	return hash
}
