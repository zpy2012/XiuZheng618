package wechat

import (
	"github.com/astaxie/beego"
	"fmt"
	"net/url"
	"github.com/astaxie/beego/httplib"
	"crypto/tls"
	"encoding/json"
	"FlowerWeChat/models"
	"FlowerWeChat/controllers/tools"
)

type WeChatAuthLoginController struct {
	beego.Controller
}

var UserId int

const (
	appid = "wx6b038c7734656d3c"
	redirect_uri = "https://flowerstore.zeaksoft.com/wechat/auth/redirect"
	response_type = "code"
	scope = "snsapi_userinfo"
	state = ""
	end = "#wechat_redirect"
	grant_type = "authorization_code"
	secret = "7e73c7a6a6aab6c89e94f7d79e7b4084"
)

func (this *WeChatAuthLoginController) Get()  {
	enCodeUrl := url.QueryEscape(redirect_uri)
	fmt.Println(enCodeUrl)
	we_chat_auth_url := fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=%s&scope=%s%s",
	appid, enCodeUrl, response_type, scope, end)
	this.Ctx.Redirect(302, we_chat_auth_url)
}

func (this *WeChatAuthLoginController) Redirect () {
	var tokenResult models.WeChatTokenResult
	var tokenCheckResult models.WeChatTokenCheckResult
	var user models.Users
	this.ParseForm(this.Ctx.Input.RequestBody)
	code := this.Ctx.Request.FormValue("code")
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=%s",
	appid, secret, code, grant_type)
	req := httplib.Get(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}else {
		json.Unmarshal([]byte(str), &tokenResult)
	}
	url = fmt.Sprintf("https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s", tokenResult.Access_token, tokenResult.Openid)
	req = httplib.Get(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	str, _ = req.String()
	json.Unmarshal([]byte(str), &tokenCheckResult)
	fmt.Println(tokenCheckResult)
	if tokenCheckResult.Code != 0 {
		url = fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s",
			appid,tokenResult.Refresh_token)
		req = httplib.Get(url)
		req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		str, _ = req.String()
		json.Unmarshal([]byte(str), &tokenResult)
	}
	url = fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN",
		tokenResult.Access_token,tokenResult.Openid)
	req = httplib.Get(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	str, _ = req.String()
	json.Unmarshal([]byte(str), &user)
	if u, err := models.GetUsersByOpenId(user.Openid); err != nil {
		user.Createdtime = tools.TimeNow().Format(beego.AppConfig.String("time"))
		id, _ := models.AddUsers(&user)
		UserId = int(id)
	}else {
		user.Id = u.Id
		models.UpdateUsersById(&user)
		UserId = user.Id
	}
	this.Ctx.Redirect(302, "/")
}
