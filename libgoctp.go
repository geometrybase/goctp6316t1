// +build linux,cgo

// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goctp

/*
#cgo linux LDFLAGS: -fPIC -L${SRCDIR}/v6.3.16_T1_20190508_api_tradeapi_se_linux64 -Wl,-rpath,${SRCDIR}/v6.3.16_T1_20190508_api_tradeapi_se_linux64 -lthostmduserapi6316t1 -lthosttraderapi6316t1 -lstdc++
#cgo linux CPPFLAGS: -fPIC -I${SRCDIR}/v6.3.16_T1_20190508_api_tradeapi_se_linux64
*/
import "C"
