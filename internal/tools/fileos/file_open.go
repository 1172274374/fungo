package fileos

import "os"

func FileWrite(path string, fileName string, fileStream []byte) bool {

	fileBuffer, err := os.OpenFile(path+fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return false
	}

	n, err := fileBuffer.Write(fileStream)
	if err != nil {
		return false
	}

	if n > 0 {
		return true
	} else {
		return false
	}
}
