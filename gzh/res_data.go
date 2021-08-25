package gzh

import "time"

/*
AccessToken公众号的全局唯一接口调用凭据
*/
type AccessTokenModel struct {
	//获取到的凭证
	AccessToken string `json:"access_token"`
	//凭证有效时间，单位：秒
	ExpiresIN time.Duration `json:"expires_in"`
}

