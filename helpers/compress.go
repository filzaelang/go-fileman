package helpers

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func CompressPdf(inputPath string, afterCompressPath string) (string, error) {
	err := api.OptimizeFile(inputPath, afterCompressPath, nil)
	if err != nil {
		return "Gagal mengompress PDF", err
	}

	return "Sukses", nil
}
