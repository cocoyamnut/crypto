# Official source snapshots

This directory contains source snapshots copied into this personal repository so they can be inspected and maintained here.

## Sources

- `go/crypto/` comes from the Go repository at commit `737d626d121e505eee448c7a1b58fa0176636b02`.
- `x-crypto/` comes from the Go supplementary crypto repository at commit `c3db4df58582058384d318c87f1c912848d8c464`.

The original source headers and upstream license files are preserved. These snapshots are source collections, not drop-in import paths: the Go standard library and `x/crypto` source use their upstream module layout and internal packages. The maintained, directly importable package in this module remains `github.com/cocoyamnut/crypto/md2` until each additional algorithm is adapted and tested under this module.

When updating an upstream snapshot, record the new commit here and review the upstream license and source changes.
