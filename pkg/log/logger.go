package pkg

import (
	"log"
	"os"
)

func RegisterLogger(msg string, filePath string) (*log.Logger, error) {
	// openLogFile, err := os.OpenFile(filePath+"/log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	return nil, err
	// }
	return log.New(os.Stdout, msg+"\t", log.Ldate|log.Ltime|log.Lshortfile), nil
}
