// Package cpu provides the small subset of golang.org/x/sys/cpu used by the
// migrated cryptographic packages in this repository.
package cpu

import (
	"github.com/cocoyamnut/crypto/internal/cpu"
	"github.com/cocoyamnut/crypto/internal/goarch"
)

const IsBigEndian = goarch.BigEndian

var X86 = cpu.X86
var Loong64 = cpu.Loong64
var S390X = cpu.S390X
