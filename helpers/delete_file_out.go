package helpers

import (
	"fmt"
	"os"
	"time"
)

func DeleteFile(filePathOut string) {
	time.Sleep(5 * time.Second)
	err := os.Remove(filePathOut)
	if err != nil {
		fmt.Println("Gagal menghapus file:", err)
	}
}
