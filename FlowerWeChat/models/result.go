package models

type Result struct {
	Code string
	Msg string
	JumpUrl string
}

type WeChatTokenResult struct {
	Access_token string 	`json:"access_token"`
	Expires_in int   	`json:"expires_in"`
	Refresh_token string	`json:"refresh_token"`
	Openid string		`json:"openid"`
	Scope string 		`json:"scope"`
}

type WeChatTokenCheckResult struct {
	Code int        `json:"errcode"`
	Msg  string        `json:"errmsg"`
}