package bloom

import (
	"github.com/zofan/go-bitset"
	"hash"
	"hash/fnv"
)

type Bloom struct {
	bitset *bitset.BitSet
	size   uint64

	hashFNV hash.Hash64

	keys int
}

func New(size uint64, keys int) *Bloom {
	return &Bloom{
		bitset: bitset.New(size),
		keys:   keys,
		size:   size,

		hashFNV: fnv.New64(),
	}
}

func (b *Bloom) Test(data []byte) bool {
	for n := 0; n < b.keys; n++ {
		if !b.bitset.Test(b.hashData(data, n) % b.size) {
			return false
		}
	}

	return true
}

func (b *Bloom) Add(data []byte) {
	for n := 0; n < b.keys; n++ {
		b.bitset.Set(b.hashData(data, n) % b.size)
	}
}

func (b *Bloom) LoadFile(file string) error {
	err := b.bitset.LoadFile(file)
	if err != nil {
		return err
	}

	b.size = b.bitset.Size()

	return nil
}

func (b *Bloom) SaveFile(file string) error {
	return b.bitset.SaveFile(file)
}

func (b *Bloom) hashData(data []byte, i int) uint64 {
	algo := b.hashFNV

	algo.Reset()
	_, _ = algo.Write(data)
	_, _ = algo.Write([]byte{
		byte(0xff & i),
		byte(0xff & (i >> 8)),
		byte(0xff & (i >> 16)),
		byte(0xff & (i >> 24)),
	})
	return algo.Sum64()
}
