package cli

import (
	"log"
	"os"
)

func ExportToFile(result, filePath string) {

	if filePath != "" {
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println("Error opening file:", err)
			return
		}
		defer file.Close()
		if _, err := file.WriteString(result); err != nil {
			log.Println("Error writing to file:", err)
		} else {
			log.Println("Saved to:", filePath)
		}
	}

}
