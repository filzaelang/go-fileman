package helpers

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func AddPDFWatermark(filePath string, filePathOut string) error {
	onTop := false
	var u types.DisplayUnit = types.POINTS
	watermarkText := "Printed by admin"
	watermarkDescription := "rot:0, op:0.5, c:0.5 0.5 0.5"
	wm, err := pdfcpu.ParseTextWatermarkDetails(watermarkText, watermarkDescription, onTop, u)
	if err != nil {
		return err
	}

	err = api.AddWatermarksFile(filePath, filePathOut, nil, wm, nil)
	if err != nil {
		return err
	}

	return nil
}
