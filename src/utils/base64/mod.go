package base64

import (
	"encoding/base64"
)

func Base64(data string) string {
	base64Data := base64.StdEncoding.EncodeToString([]byte(data))
	// return base64Data
	return base64Data
}
