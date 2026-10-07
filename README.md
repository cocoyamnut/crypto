# md2

`github.com/cocoyamnut/crypto/md2` is a small, standard-library-style Go implementation of the MD2 message-digest algorithm from [RFC 1319](https://www.rfc-editor.org/rfc/rfc1319).

The package is organized like the Go MD4/MD5 implementations: `md2.go` contains the public API and streaming state, while `md2block.go` contains the block transformation.

MD2 is obsolete and cryptographically broken. Use this package only for legacy compatibility, protocol interoperability, or historical data.

## API

```go
import "github.com/cocoyamnut/crypto/md2"

h := md2.New()
_, _ = h.Write([]byte("hello"))
digest := h.Sum(nil)

oneShot := md2.Sum([]byte("hello")) // [16]byte
```

The package exposes `New() hash.Hash`, `Sum([]byte) [16]byte`, `Size` (16), and `BlockSize` (16). The test suite includes every RFC 1319 test vector and behavior tests for chunked writes, `Sum`, and `Reset`.

## License

MIT. See [LICENSE](LICENSE).
