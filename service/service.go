package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func GetRequestParams(r *http.Request, requestParams interface{}) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）

	contentType := r.Header.Get("Content-type")

	if contentType == "｀application/json" {
		json.NewDecoder(r.Body).Decode(&requestParams)
	} else if strings.Index(contentType, "application/x-www-form-urlencoded") != -1 {
		rp := make(map[string]interface{})
		for i, num := range r.Form {
			// println("i:", i, " num:", num[0])
			rp[i] = num[0]
		}

		rpJson, _ := json.Marshal(rp)
		fmt.Println("requestParams:", string(rpJson))

		// reader := strings.NewReader(string(rpJson))
		// json.NewDecoder(reader).Decode(&requestParams)

		json.Unmarshal(rpJson, &requestParams)
	}
}

// CheckFileExist /**
func CheckFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
