package GoNet

import "net/url"

func URLEncode(str string) string {
	result := url.QueryEscape(str)
	return result
}

func URLDecode(str string) string {
	result, _ := url.QueryUnescape(str)
	return result
}
