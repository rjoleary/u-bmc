// Copyright 2018-2019 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !arm
// Assume non arm host is the host system

package aspeed

func openMem() memProvider {
	return openLpcMemory(0x2e)
}
