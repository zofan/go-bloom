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
		{``, 37},
		{`hello`, 11},
		{`world`, 19},
		{`1000`, 36},
		{`10000`, 60},
		{`100000`, 36},
		{`golang`, 35},
	}

	for _, c := range cases {
		if bn := hash([]byte(c.key), 1) % 64; bn != c.bitNum {
			t.Errorf(`expected bit number %d for key %s, given %d`, c.bitNum, c.key, bn)
		}
	}
}

func TestA(t *testing.T) {
	bs := bitset.New(64)
	b := New(bs, 3)

	err := b.Add([]byte(`hello`))
	if err != nil {
		t.Error(`add key, expected empty error`)
	}

	if !b.Test([]byte(`hello`)) {
		t.Error(`test key, expected true`)
	}

	if b.Test([]byte(`golang'`)) {
		t.Error(`test key, expected false`)
	}
}

func TestFullCollisions(t *testing.T) {
	bs := bitset.New(64)
	b := New(bs, 128)

	err := b.Add([]byte(`1000`))
	if err != nil {
		t.Error(`add key, expected empty error`)
	}

	if !b.Test([]byte(`1000`)) {
		t.Error(`test key, expected true`)
	}

	if !b.Test([]byte(`100000`)) {
		t.Error(`test key, expected true`)
	}

	if !b.Test([]byte(`hello`)) {
		t.Error(`test key, expected true`)
	}
}

func TestPartialCollision(t *testing.T) {
	bs := bitset.New(64)
	b := New(bs, 3)

	err := b.Add([]byte(`hello`))
	if err != nil {
		t.Error(`add key, expected empty error`)
	}

	if !b.Test([]byte(`world`)) {
		t.Error(`test key, expected true`)
	}

	if b.Test([]byte(`golang`)) {
		t.Error(`test key, expected false`)
	}
}
