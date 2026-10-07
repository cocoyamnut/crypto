// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gcm_test

import (
	"github.com/cocoyamnut/crypto/cipher"
	"github.com/cocoyamnut/crypto/internal/crypto/fips140/aes/gcm"
)

var _ cipher.AEAD = (*gcm.GCM)(nil)