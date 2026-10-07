# Algorithms at repository root

All collected algorithm source directories live at the repository root beside `go.mod`. The directory names identify the algorithm or primitive; `_xcrypto` marks a second official implementation with the same package name as a standard-library directory.

`md2/` is the implementation maintained directly in this repository. The other root-level directories are source snapshots moved from the Go standard library and `golang.org/x/crypto`; their original headers, tests, and licenses are preserved.

The snapshots are intentionally kept together for personal maintenance. They may still require import-path and internal-package adaptation before being used as ordinary packages under this module.
