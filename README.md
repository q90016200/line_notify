[TOC]

## 環境變數 .env
```
PORT=9123
```
>有引入 [line_notify_package](https://github.com/q90016200/line_notify_package), env 需添加
```
LINE_NOTIFY_CLIENT_ID=
LINE_NOTIFY_CLIENT_SECRET=
LINE_NOTIFY_CALLBACK_URL=http://127.0.0.1:9123/lineNotify/callback
LINE_NOTIFY_TOKEN_FILE=lineNotifyAccessToken.txt
```

若是本地網址為 http://127.0.0.1:9123

## 使用方法
### 保存 accessToken 
進入 `/lineNotify/auth` 頁面內，選擇需要傳送的對象後，即會在 env `LINE_NOTIFY_TOKEN_FILE` 保存 access_token

### send notify
URL : /lineNotify/notify

METHOD: POST

| 參數名 | 類型 | 必填 | 描述 |
| ------- | -------- | --- | --------------------------- |
| message | string | Y | 訊息 |
| access_token | string | N | 發送 LINE 通知 需要的 token ，若不填寫會使用 auth 保存的 access_token 進行使用 |

