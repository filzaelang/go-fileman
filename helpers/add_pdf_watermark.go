package helpers

import (
	"time"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func AddPDFWatermark(filePath string, filePathOut string, username string) error {
	onTop := false
	var u types.DisplayUnit = types.POINTS
	watermarkText := "Printed by " + username + "\n" + time.Now().Format("02 Jan 2006 15:04")
	watermarkDescription := "font:Helvetica-Oblique, points:12, c:0 0 1, scale:1 abs, rot:0, pos:br, off:-15 15"
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
