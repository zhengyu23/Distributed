package utils

import (
	"net/http"
	"strconv"
)

func GetHashFromHeader(h http.Header) string {
	digest := h.Get("Digest")
	if len(digest) < 9 {
		return ""
	}
	if digest[:8] != "SHA-256=" {
		return ""
	}
	return digest[8:]
}
func GetSizeFromHeader(h http.Header) int64 {
	size, _ := strconv.ParseInt(h.Get("Content-Length"), 0, 64)
	return size
}
