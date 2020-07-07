package api

import (
	"encoding/json"
	"lineNotify/service"
	"lineNotify/service/lineNotify"
	"net/http"
)

type CallbackRequestStuct struct {
	Code string `json:"code"`
}

type 

func LineNotifyAuth(w http.ResponseWriter, r *http.Request) {
	// http.Redirect(w, r, lineNotify.Auth("none"), http.StatusSeeOther)
	data := make(map[string]interface{})
	err := make(map[string]interface{})

	data["redirect"] = lineNotify.Auth("none")
	Response(w, data, err)
}

func LineNotifyCallback(w http.ResponseWriter, r *http.Request) {
	// 處理請求
	requestParams := CallbackRequestStuct{}
	service.GetRequestParams(r, &requestParams)

	// 取得 code
	oauth = lineNotify.OauthToken(requestParams.Code)

}

func LineNotifySendMessage(w http.ResponseWriter, r *http.Request) {

}

type Res struct {
	Data *interface{} `json:"data"`
	Err  *interface{} `json:"error"`
}

func Response(w http.ResponseWriter, data interface{}, err interface{}) {
	res := Res{
		Data: &data,
		Err:  &err,
	}

	j, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(200)
	w.Write(j)
}
