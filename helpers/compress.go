package helpers

import (
	"os"

	"fmt"
	"image/png"
	"os/exec"
)

func CompressPdf(srcPath, dstPath string) error {
	cmd := exec.Command("gswin64c",
		"-sDEVICE=pdfwrite",
		"-dCompatibilityLevel=1.4",
		"-dPDFSETTINGS=/ebook",
		"-dAutoRotatePages=/None",
		"-dNOPAUSE",
		"-dBATCH",
		fmt.Sprintf("-sOutputFile=%s", dstPath),
		srcPath,
	)

	_, err := cmd.CombinedOutput()
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

// func CompressPdf(srcPath, dstPath string) error {
// 	conf := model.NewDefaultConfiguration()

// 	conf.Optimize = true
// 	conf.OptimizeBeforeWriting = true
// 	conf.OptimizeResourceDicts = true
// 	conf.OptimizeDuplicateContentStreams = true

// 	err := api.OptimizeFile(srcPath, dstPath, conf)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
