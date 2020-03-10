package imaging

import (
	"image"
	"testing"
)

func TestResize(t *testing.T) {
	td := []struct {
		desc string
		src  image.Image
		w, h int
		f    ResampleFilter
		want *image.NRGBA
	}{
		{
			"Resize 2x2 1x1 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			1, 1,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix:    []uint8{0x40, 0x40, 0x40, 0xc0},
			},
		},
		{
			"Resize 2x2 2x2 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			2, 2,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
		},
		{
			"Resize 3x1 1x1 nearest",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 0),
				Stride: 3 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			1, 1,
			NearestNeighbor,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix:    []uint8{0x00, 0xff, 0x00, 0xff},
			},
		},
		{
			"Resize 2x2 0x4 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			0, 4,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 4, 4),
				Stride: 4 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
		},
		{
			"Resize 2x2 4x0 linear",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			4, 0,
			Linear,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 4, 4),
				Stride: 4 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x40, 0xbf, 0x00, 0x00, 0xbf, 0xff, 0x00, 0x00, 0xff,
					0x00, 0x40, 0x00, 0x40, 0x30, 0x30, 0x10, 0x70, 0x8f, 0x10, 0x30, 0xcf, 0xbf, 0x00, 0x40, 0xff,
					0x00, 0xbf, 0x00, 0xbf, 0x10, 0x8f, 0x30, 0xcf, 0x30, 0x30, 0x8f, 0xef, 0x40, 0x00, 0xbf, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0xbf, 0x40, 0xff, 0x00, 0x40, 0xbf, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
		},
	}
	for _, d := range td {
		got := Resize(d.src, d.w, d.h, d.f)
		want := d.want
		if !compareNRGBA(got, want, 1) {
			t.Errorf("test [%s] failed: %#v", d.desc, got)
		}
	}
}

func TestFit(t *testing.T) {
	td := []struct {
		desc string
		src  image.Image
		w, h int
		f    ResampleFilter
		want *image.NRGBA
	}{
		{
			"Fit 2x2 1x10 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			1, 10,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix:    []uint8{0x40, 0x40, 0x40, 0xc0},
			},
		},
		{
			"Fit 2x2 10x1 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			10, 1,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix:    []uint8{0x40, 0x40, 0x40, 0xc0},
			},
		},
		{
			"Fit 2x2 10x10 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			10, 10,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
		},
	}
	for _, d := range td {
		got := Fit(d.src, d.w, d.h, d.f)
		want := d.want
		if !compareNRGBA(got, want, 0) {
			t.Errorf("test [%s] failed: %#v", d.desc, got)
		}
	}
}

func TestThumbnail(t *testing.T) {
	td := []struct {
		desc string
		src  image.Image
		w, h int
		f    ResampleFilter
		want *image.NRGBA
	}{
		{
			"Thumbnail 6x2 1x1 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 5, 1),
				Stride: 6 * 4,
				Pix: []uint8{
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			1, 1,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix:    []uint8{0x40, 0x40, 0x40, 0xc0},
			},
		},
		{
			"Thumbnail 2x6 1x1 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 5),
				Stride: 2 * 4,
				Pix: []uint8{
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			1, 1,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix:    []uint8{0x40, 0x40, 0x40, 0xc0},
			},
		},
		{
			"Thumbnail 1x3 2x2 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 0, 2),
				Stride: 1 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00,
					0xff, 0x00, 0x00, 0xff,
					0xff, 0xff, 0xff, 0xff,
				},
			},
			2, 2,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
				},
			},
		},
	}
	for _, d := range td {
		got := Thumbnail(d.src, d.w, d.h, d.f)
		want := d.want
		if !compareNRGBA(got, want, 0) {
			t.Errorf("test [%s] failed: %#v", d.desc, got)
		}
	}
}
