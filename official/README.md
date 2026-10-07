# Official Go cryptography index

This repository collects personal implementations and compatibility packages. Official Go implementations remain the source of truth and are used directly where available; they are not copied into this module.

## Standard library (`crypto/...`)

| Package | Algorithms / purpose |
|---|---|
| `crypto/aes` | AES / Rijndael |
| `crypto/cipher` | Block-cipher modes and interfaces |
| `crypto/des` | DES and 3DES/TDEA |
| `crypto/dsa` | DSA |
| `crypto/ecdh` | ECDH, including X25519 |
| `crypto/ecdsa` | ECDSA |
| `crypto/ed25519` | Ed25519 |
| `crypto/elliptic` | NIST elliptic curves |
| `crypto/hmac` | HMAC |
| `crypto/md5` | MD5 |
| `crypto/rand` | Cryptographically secure randomness |
| `crypto/rc4` | RC4 |
| `crypto/rsa` | RSA |
| `crypto/sha1` | SHA-1 |
| `crypto/sha256` | SHA-224 and SHA-256 |
| `crypto/sha3` | SHA-3 and SHAKE |
| `crypto/sha512` | SHA-384, SHA-512/224, SHA-512/256, SHA-512 |
| `crypto/subtle` | Constant-time helpers |
| `crypto/tls` | TLS 1.2/1.3 |
| `crypto/x509` | X.509 and PKIX |

## Official supplementary module (`golang.org/x/crypto/...`)

| Package | Algorithms / purpose |
|---|---|
| `argon2` | Argon2 password KDF |
| `bcrypt` | bcrypt password hashing |
| `blake2b`, `blake2s` | BLAKE2 |
| `chacha20`, `chacha20poly1305` | ChaCha20 and ChaCha20-Poly1305 |
| `curve25519` | X25519 |
| `ed25519` | Ed25519 support package |
| `hkdf` | HKDF |
| `md4` | MD4 |
| `poly1305` | Poly1305 MAC |
| `ripemd160` | RIPEMD-160 |
| `sha3` | SHA-3, SHAKE, and legacy Keccak |
| `twofish` | Twofish |

The `md2` package in this repository fills a gap not provided by either official collection:

```go
import "github.com/cocoyamnut/crypto/md2"
```

For official implementations, import the upstream package directly so security fixes and architecture-specific improvements can be received normally.
