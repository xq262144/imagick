// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#cgo LDFLAGS: -ljpeg -lpng -lwebp -lz -lgomp -lm -lbz2
#cgo !no_pkgconfig pkg-config: MagickWand MagickCore
*/
import "C"
