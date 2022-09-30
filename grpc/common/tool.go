package common

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
)

func GetIP() (string, error) {
	res, err := http.Get("http://whatismyip.akamai.com")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	ip := string(body)
	return ip, nil
}

// CalcMD5 计算文件 MD5
func CalcMD5(fileData []byte) (string, error) {
	md5hash := md5.New()
	_, err := md5hash.Write(fileData)
	if err != nil {
		return "", err
	}
	md5Str := hex.EncodeToString(md5hash.Sum(nil))
	return md5Str, nil
}
