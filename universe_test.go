// Copyright 2021 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pps

import (
	"image/color"
	"testing"
)

func TestParticleColor(t *testing.T) {
	tests := []struct {
		N      int
		NClose int
		want   color.Color
	}{
		{N: 0, want: colorGreen},
		{N: 1, want: colorGreen},
		{N: 12, want: colorGreen},
		{N: 13, want: colorBrown},
		{N: 15, want: colorBrown},
		{N: 16, want: colorBlue},
		{N: 35, want: colorBlue},
		{N: 36, want: colorYellow},
		{N: 50, want: colorYellow},
		{N: 50, NClose: 14, want: colorYellow},
		{N: 20, NClose: 16, want: colorMagenta},
		{N: 20, NClose: 17, want: colorMagenta},
		{N: 35, NClose: 16, want: colorMagenta},
		{N: 50, NClose: 16, want: colorMagenta},
	}
	for _, tt := range tests {
		got := particleColor(tt.N, tt.NClose)
		if got != tt.want {
			t.Errorf("particleColor(N: %d, NClose: %d) = %v, want %v",
				tt.N, tt.NClose, got, tt.want)
		}
	}
}
