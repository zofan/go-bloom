package bloom

import (
	"github.com/zofan/go-bitset"
	"testing"
)

func TestHash(t *testing.T) {
	cases := []struct {
		key    string
		bitNum uint64
	}{
		{``, 0},
		{`hello`, 12},
		{`world`, 12},
		{`1000`, 63},
		{`2000`, 17},
		{`10000`, 5},
		{`100000`, 18},
		{`golang`, 38},
		{"\n\n\n", 48},
		{"\t\t\t", 43},
	}

	b := New(bitset.New(64), 1)

	for i, c := range cases {
		if bn := b.hashData([]byte(c.key), i) % 64; bn != c.bitNum {
			t.Errorf(`expected bit number %d for key %s, given %d`, c.bitNum, c.key, bn)
		}
	}
}

func TestA(t *testing.T) {
	bs := bitset.New(64)
	b := New(bs, 3)

	b.Add([]byte(`hello`))

	if !b.Test([]byte(`hello`)) {
		t.Error(`test key, expected true`)
	}

	if b.Test([]byte(`golang'`)) {
		t.Error(`test key, expected false`)
	}
}
