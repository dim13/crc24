package crc24

import "hash"

const (
	initial = 0x0b704ce
	polyhom = 0x1864cfb
)

type digest struct {
	sum uint32
}

func (d *digest) Write(p []byte) (n int, err error) {
	for _, v := range p {
		d.sum ^= uint32(v) << 16
		for i := 0; i < 8; i++ {
			d.sum <<= 1
			if d.sum&0x1000000 != 0 {
				d.sum ^= polyhom
			}
		}
	}
	return len(p), nil
}

func (d *digest) Sum(b []byte) []byte {
	v := d.Sum32()
	for i := d.Size() - 1; i >= 0; i-- {
		b = append(b, byte(v>>uint(8*i)))
	}
	return b
}

func (d *digest) Reset() {
	d.sum = initial
}

func (d *digest) Size() int {
	return 4
}

func (d *digest) BlockSize() int {
	return 1
}

func (d *digest) Sum32() uint32 {
	return d.sum & 0xffffff
}

func New() hash.Hash32 {
	d := &digest{}
	d.Reset()
	return d
}

func Sum(b []byte) uint32 {
	d := New()
	d.Write(b)
	return d.Sum32()
}

func SumString(s string) uint32 {
	return Sum([]byte(s))
}
