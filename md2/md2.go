// Package md2 implements the MD2 message-digest algorithm specified by RFC 1319.
//
// MD2 is provided for compatibility with legacy protocols and data. It is
// cryptographically broken and must not be used for new security designs.
package md2

import (
	"hash"

	"github.com/cocoyamnut/crypto"
)

const (
	// Size is the length of an MD2 checksum in bytes.
	Size = 16
	// BlockSize is the MD2 input block size in bytes.
	BlockSize = 16
)

var _ hash.Hash = (*digest)(nil)

func init() {
	crypto.RegisterHash(crypto.MD2, New)
}

type digest struct {
	state    [Size]byte
	checksum [Size]byte
	buffer   [BlockSize]byte
	n        int
}

// New returns a new MD2 hash.Hash.
func New() hash.Hash {
	d := new(digest)
	d.Reset()
	return d
}

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
