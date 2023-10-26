package fileos

import (
	"bytes"
	"mime/multipart"
	"path"
)

// FileParse 解析文件
func FileParse(file multipart.File, fileName string) ([]byte, string, string) {

	//文件前缀
	filePrefix := path.Base(fileName)
	//文件后缀
	fileSuffix := path.Ext(fileName)
	// 将file转换成字节流
	fil := make([][]byte, 0)
	var b int64 = 0
	// 通过for循环写入
	for {
		buffer := make([]byte, 1024)
		n, err := file.ReadAt(buffer, b)
		b = b + int64(n)
		fil = append(fil, buffer)
		if err != nil {
			break
		}
	}

	fileStream := bytes.Join(fil, []byte(""))
	return fileStream, filePrefix, fileSuffix
}
