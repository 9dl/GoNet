package GoNet

import (
	"encoding/json"
	"strings"
	"time"
)

func Replace(input string, old string, new string) string {
	return strings.Replace(input, old, new, -1)
}

func CountStringOccurrences(input string, substr string) int {
	return strings.Count(input, substr)
}

func CurrentUnixTime() int64 {
	return time.Now().Unix()
}

func ParseJson(jsonStr string) map[string]interface{} {
	var result map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &result)
	return result
}
