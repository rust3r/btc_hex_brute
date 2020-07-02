package brute

import (
	"log"
	"os"
)

func saveData(data string, path string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0655)
	if err != nil {
		log.Fatal(data, err)
	}
	defer f.Close()

	_, err = f.WriteString(data)
	if err != nil {
		log.Fatal(err)
	}
}
