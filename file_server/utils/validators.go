package utils

import (
	"mime/multipart"
	"net/http"
)

func IsValidStructure(file multipart.File, extentions []string) bool {
	buffer := make([]byte, 512)

	file.Read(buffer)

	file.Seek(0, 0)

	mimetype := http.DetectContentType(buffer)

	for _, t := range extentions {
		if mimetype == t {
			return true
		}
	}

	return false
}

func IsValidImage(file multipart.File) bool {
	return IsValidStructure(file, []string{"image/jpeg", "image/png", "image/gif", "image/webp"})
}

func IsValidVideo(file multipart.File) bool {
	return IsValidStructure(file, []string{"video/mp4"})
}
