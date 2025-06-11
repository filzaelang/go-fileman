package helpers

import (
	"regexp"
	"strings"
)

func PrepFilename(filename string) string {
	extPos := strings.LastIndex(filename, ".")
	if extPos == -1 {
		return filename
	}

	ext := filename[extPos:]
	name := filename[:extPos]
	safeName := strings.ReplaceAll(name, ".", "_")

	return safeName + ext
}

func NormalizeFilename(fileraw string) string {
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(fileraw, "_")
}

// function _prep_filename($filename)
//     {
//         if (($ext_pos = strrpos($filename, '.')) === FALSE) {
//             return $filename;
//         }

//         $ext = substr($filename, $ext_pos);
//         $filename = substr($filename, 0, $ext_pos);
//         return str_replace('.', '_', $filename) . $ext;
//     }
