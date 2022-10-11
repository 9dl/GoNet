package GoNet

import "encoding/base64"

func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64Decode(str string) string {
	data, _ := base64.StdEncoding.DecodeString(str)
	return string(data)
}
