package coding

import (
	"encoding/base64"
	"io/ioutil"
)

// 从图片文件获取base64编码
func FileToBase64Code(dst string) (string, error) {
	buf, err := ioutil.ReadFile(dst)
	return base64.StdEncoding.EncodeToString(buf), err
}

// base64图片编码保存到文件
func Base64CodeToFile(code, dst string) error {
	buf, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dst, buf, 0640)
}
