package api

import (
	"encoding/json"
	"fmt"
	"html/template"
	"lineNotify/service"
	"net/http"
	"os"

	"github.com/q90016200/line_notify_package/lineNotify"
)

type CallbackRequestStuct struct {
	Code string `json:"code"`
}

func LineNotifyAuth(w http.ResponseWriter, r *http.Request) {
	// http.Redirect(w, r, lineNotify.Auth("none"), http.StatusSeeOther)
	data := make(map[string]interface{})
	// err := make(map[string]interface{})

	data["redirect"] = lineNotify.Auth("none")
	Response(w, 200, data)
}

func LineNotifyCallback(w http.ResponseWriter, r *http.Request) {
	// 處理請求
	requestParams := CallbackRequestStuct{}
	service.GetRequestParams(r, &requestParams)

	// 撤銷 access token
	ln := lineNotify.NewLineNotify("")
	ln.Revoke()

	// 取得 code
	oauth := lineNotify.OauthToken(requestParams.Code)
	if oauth != "none" {
		// 寫入 file lineNotifyAccessToken.txt
		fileName := os.Getenv("LINE_NOTIFY_TOKEN_FILE")
		f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("write")
			f.Write([]byte(oauth))
		}
		f.Close()
	}

	fmt.Println("[oauth] accessToken:", oauth)
}

type SendNotifyRequestParams struct {
	Message string `json:"message"`
	AccessToken string `json:"access_token"`
}

func LineNotifySendNotify(w http.ResponseWriter, r *http.Request) {
	// 處理請求
	requestParams := SendNotifyRequestParams{}
	service.GetRequestParams(r, &requestParams)

	accessToken := ""
	if requestParams.AccessToken != "" {
		accessToken = requestParams.AccessToken
	}

	ln := lineNotify.NewLineNotify(accessToken)
	notify := ln.Notify(requestParams.Message)

	data := make(map[string]interface{})
	data["message"] = requestParams.Message

	if !notify {
		Response(w, 400, data)
	} else {
		Response(w, 200, data)
	}

}

func LineNotifyIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/index.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	fmt.Println("LineNotifyIndex")
	t.Execute(w, "data goes here")
}

type Res struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func Response(w http.ResponseWriter, status int, data interface{}) {
	res := Res{
		Status: status,
		Data:   data,
	}

	j, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(status)
	w.Write(j)
}
