// Package md2 implements the MD2 message-digest algorithm specified by RFC 1319.
//
// MD2 is provided for compatibility with legacy protocols and data. It is
// cryptographically broken and must not be used for new security designs.
package md2

import "hash"

const (
	// Size is the length of an MD2 checksum in bytes.
	Size = 16
	// BlockSize is the MD2 input block size in bytes.
	BlockSize = 16
)

var _ hash.Hash = (*digest)(nil)

type digest struct {
	state    [Size]byte
	checksum [Size]byte
	buffer   [BlockSize]byte
	n        int
}

// New returns a new MD2 hash.Hash.
func New() hash.Hash { return new(digest) }

// Sum returns the MD2 checksum of data.
func Sum(data []byte) [Size]byte {
	d := digest{}
	_, _ = d.Write(data)
	return d.checksumFinal()
}

func (d *digest) Write(p []byte) (int, error) {
	original := len(p)
	if d.n != 0 {
		copied := copy(d.buffer[d.n:], p)
		d.n += copied
		p = p[copied:]
		if d.n == BlockSize {
			d.process(d.buffer[:])
			d.n = 0
		}
	}
	for len(p) >= BlockSize {
		d.process(p[:BlockSize])
		p = p[BlockSize:]
	}
	if len(p) != 0 {
		d.n = copy(d.buffer[:], p)
	}
	return original, nil
}

func (d *digest) Sum(in []byte) []byte {
	clone := *d
	sum := clone.checksumFinal()
	return append(in, sum[:]...)
}

func (d *digest) Reset() {
	*d = digest{}
}

func (d *digest) Size() int      { return Size }
func (d *digest) BlockSize() int { return BlockSize }

func (d *digest) checksumFinal() [Size]byte {
	padLen := BlockSize - d.n
	var finalBlock [BlockSize]byte
	copy(finalBlock[:], d.buffer[:d.n])
	for i := d.n; i < BlockSize; i++ {
		finalBlock[i] = byte(padLen)
	}
	d.process(finalBlock[:])
	d.process(d.checksum[:])
	return d.state
}

func (d *digest) process(block []byte) {
	var x [48]byte
	copy(x[:Size], d.state[:])
	copy(x[Size:2*Size], block)
	for i := 0; i < Size; i++ {
		x[2*Size+i] = d.state[i] ^ block[i]
	}

	t := byte(0)
	for j := 0; j < 18; j++ {
		for k := range x {
			x[k] ^= sBox[t]
			t = x[k]
		}
		t += byte(j)
	}
	copy(d.state[:], x[:Size])

	last := d.checksum[Size-1]
	for i := 0; i < Size; i++ {
		d.checksum[i] ^= sBox[block[i]^last]
		last = d.checksum[i]
	}
}

var sBox = [256]byte{
	41, 46, 67, 201, 162, 216, 124, 1, 61, 54, 84, 161, 236, 240, 6, 19,
	98, 167, 5, 243, 192, 199, 115, 140, 152, 147, 43, 217, 188, 76, 130, 202,
	30, 155, 87, 60, 253, 212, 224, 22, 103, 66, 111, 24, 138, 23, 229, 18,
	190, 78, 196, 214, 218, 158, 222, 73, 160, 251, 245, 142, 187, 47, 238, 122,
	169, 104, 121, 145, 21, 178, 7, 63, 148, 194, 16, 137, 11, 34, 95, 33,
	128, 127, 93, 154, 90, 144, 50, 39, 53, 62, 204, 231, 191, 247, 151, 3,
	255, 25, 48, 179, 72, 165, 181, 209, 215, 94, 146, 42, 172, 86, 170, 198,
	79, 184, 56, 210, 150, 164, 125, 182, 118, 252, 107, 226, 156, 116, 4, 241,
	69, 157, 112, 89, 100, 113, 135, 32, 134, 91, 207, 101, 230, 45, 168, 2,
	27, 96, 37, 173, 174, 176, 185, 246, 28, 70, 97, 105, 52, 64, 126, 15,
	85, 71, 163, 35, 221, 81, 175, 58, 195, 92, 249, 206, 186, 197, 234, 38,
	44, 83, 13, 110, 133, 40, 132, 9, 211, 223, 205, 244, 65, 129, 77, 82,
	106, 220, 55, 200, 108, 193, 171, 250, 36, 225, 123, 8, 12, 189, 177, 74,
	120, 136, 149, 139, 227, 99, 232, 109, 233, 203, 213, 254, 59, 0, 29, 57,
	242, 239, 183, 14, 102, 88, 208, 228, 166, 119, 114, 248, 235, 117, 75, 10,
	49, 68, 80, 180, 143, 237, 31, 26, 219, 153, 141, 51, 159, 17, 131, 20,
}
