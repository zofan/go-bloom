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
		{`hello`, 7},
		{`world`, 7},
		{`golang`, 41},
	}

	for _, c := range cases {
		if bn := hash([]byte(c.key), 0) % 64; bn != c.bitNum {
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
	b := New(bs, 64)

	err := b.Add([]byte(`hello`))
	if err != nil {
		t.Error(`add key, expected empty error`)
	}

	if !b.Test([]byte(`hello`)) {
		t.Error(`test key, expected true`)
	}

	if !b.Test([]byte(`world`)) {
		t.Error(`test key, expected true`)
	}

	if !b.Test([]byte(`golang`)) {
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
