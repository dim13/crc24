package crc24

import "testing"

func TestSumString(t *testing.T) {
	testCases := []struct {
		s string
		n uint32
	}{
		{"", 0xb704ce},
		{"test", 0xf86ed0},
		{"aaaaaaaaaaaaaaaaaaaa", 0xbca062},
	}
	for _, tc := range testCases {
		t.Run(tc.s, func(t *testing.T) {
			sum := SumString(tc.s)
			if sum != tc.n {
				t.Errorf("got %v, want %v", sum, tc.n)
			}
		})
	}
}

func BenchmarkWrite(b *testing.B) {
	d := New()
	p := []byte("the quick brown fox jumps over the lazy dog")
	for i := 0; i < b.N; i++ {
		d.Write(p)
	}
}
