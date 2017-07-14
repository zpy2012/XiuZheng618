package wechat

import (
	"github.com/astaxie/beego"
	"sort"
	"crypto/sha1"
	"io"
	"strings"
	"fmt"
)

type WeChatAuthController struct {
	beego.Controller
}

const (
	token = "flowerstore"
)

func makeSignature(timestamp, nonce string) string {
	sl := []string{token, timestamp, nonce}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}

func (this *WeChatAuthController)validateUrl() (success bool, str string) {
	timestamp := strings.Join(this.Ctx.Request.Form["timestamp"], "")
	nonce := strings.Join(this.Ctx.Request.Form["nonce"], "")
	signatureGen := makeSignature(timestamp, nonce)
	fmt.Println("4444444444", timestamp, nonce, signatureGen)
	signatureIn := strings.Join(this.Ctx.Request.Form["signature"], "")
	if signatureGen != signatureIn {
		return false, ""
	}
	echostr := strings.Join(this.Ctx.Request.Form["echostr"], "")
	return true, echostr
}

func (this *WeChatAuthController)ProcRequest() {
	this.ParseForm(this.Ctx.Input.RequestBody)
	if success, echostr := this.validateUrl();success == true {
		fmt.Fprint(this.Ctx.ResponseWriter, echostr)
	}else {
		fmt.Fprint(this.Ctx.ResponseWriter, "false")
	}
}

func (this *WeChatAuthController) File () {
	fmt.Fprint(this.Ctx.ResponseWriter, "G5nTSndQ8eFMpvZJ")
}