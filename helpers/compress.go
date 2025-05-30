package helpers

import (
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"

	"image/png"
)

func CompressPdf(srcPath, dstPath string) error {
	err := api.OptimizeFile(srcPath, dstPath, nil)
	if err != nil {
		return err
	}

	return nil
}

func CompressPng(srcPath, dstPath string) error {
	f, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer f.Close()

	img, e := png.Decode(f)
	if e != nil {
		return e
	}

	// Buat file tujuan
	out, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Encoder with maximum compression
	encoder := png.Encoder{
		CompressionLevel: png.BestCompression,
	}

	return encoder.Encode(out, img)
}
