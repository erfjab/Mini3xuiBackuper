package utils

import (
	"fmt"
	"time"
)

func BackupFilename(username string) string {
	return fmt.Sprintf("%s_%s.db", time.Now().Format("2006-01-02_15-04-05"), username)
}
