package lineNotify

import (
	"net/http"
	"net/url"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Notify struct {
	accessToken string
}

type OauthTokenResponseStruct struct {
	access_token string
}

// 用來建構 Notify 的假建構子
func LineNotify(accessToken string) (notify *Notify) {
	notify = &Notify{accessToken: accessToken}

	// 這裡會回傳一個型態是 *Notify 建構體的 notify 變數
	return notify
}

func Auth(state string) string {
	clientID := os.Getenv("LINE_NOTIFY_CLIENT_ID")
	// clientSecret := os.Getenv("LINE_NOTIFY_CLIENT_SECRET")
	callbackURL := os.Getenv("LINE_NOTIFY_CALLBACK_URL")

	return "https://notify-bot.line.me/oauth/authorize?response_type=code&scope=notify&response_mode=form_post&client_id=" + clientID + "&redirect_uri=" + callbackURL + "&state=" + state
}

func OauthToken(code string) bool {
	// get access_token
	postURL := "https://notify-bot.line.me/oauth/token"
	postParams := url.Values{}
	postParams.Add("grant_type", "authorization_code")
	postParams.Add("code", code)
	postParams.Add("redirect_uri", os.Getenv("LINE_NOTIFY_CALLBACK_URL"))
	postParams.Add("client_id", os.Getenv("LINE_NOTIFY_CLIENT_ID"))
	postParams.Add("client_secret", os.Getenv("LINE_NOTIFY_CLIENT_SECRET"))

	resp, err := http.PostForm(postURL, postParams)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()

	return false
}
