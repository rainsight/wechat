// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package thirdparty

import (
	"net/url"
)

// 请求用户授权时跳转的地址.
func AuthCodeURL(suiteId, preAuthCode, redirectURI, state string) string {
	// https://qy.weixin.qq.com/cgi-bin/loginpage?suite_id=$suite_id$
	// &pre_auth_code=$pre_auth_code$&redirect_uri=$redirect_uri$&state=$state$
	return "https://qy.weixin.qq.com/cgi-bin/loginpage?suite_id=" + url.QueryEscape(suiteId) +
		"&pre_auth_code=" + url.QueryEscape(preAuthCode) +
		"&redirect_uri=" + url.QueryEscape(redirectURI) +
		"&state=" + url.QueryEscape(state)
}
