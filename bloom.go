package bloom

import (
	"github.com/zofan/go-bitset"
	"hash"
	"hash/crc64"
	"hash/fnv"
)

type Bloom struct {
	bitset *bitset.BitSet

	hashFNV   hash.Hash64
	hashCrc64 hash.Hash64

	keys int
}

func New(bs *bitset.BitSet, keys int) *Bloom {
	return &Bloom{
		bitset: bs,
		keys:   keys,

		hashFNV:   fnv.New64(),
		hashCrc64: crc64.New(crc64.MakeTable(crc64.ISO)),
	}
}

func (b *Bloom) Test(data []byte) bool {
	for n := 0; n < b.keys; n++ {
		if !b.bitset.Test(b.hashData(data, n) % b.bitset.Size()) {
			return false
		}
	}

	return true
}

func (b *Bloom) Add(data []byte) {
	for n := 0; n < b.keys; n++ {
		b.bitset.Set(b.hashData(data, n) % b.bitset.Size())
	}
}

func (b *Bloom) hashData(data []byte, i int) uint64 {
	algo := b.hashFNV

	if i%2 == 0 {
		algo = b.hashCrc64
	}

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
