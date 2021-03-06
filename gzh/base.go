package gzh

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"github.com/qnsoft/go-wx-sdk/utils"
	"strings"
	"time"
)

/*---------------------------------------------------基础支持--------------------------------------------------------*/
/*
AccessToken公众号的全局唯一接口调用凭据
*/
type AccessTokenModel struct {
	//获取到的凭证
	AccessToken string `json:"access_token"`
	//凭证有效时间，单位：秒
	ExpiresIN time.Duration `json:"expires_in"`
}

/*
公众号对象
*/
type GzhApi struct {
	//appid 从公众号开发者中心获取
	AppID string `json:"appid"`
	//secret 从公众号开发者中心获取
	Secret string `json:"secret"`
}

/**
 * @Description: 获取token
 * @return *AccessTokenModel
 */
func (this *GzhApi) get_token() AccessTokenModel {
	glog.SetPath("./tmp/wx_gzh_log")
	_accessToken := AccessTokenModel{}
	// 获取缓存
	_token_info, _ := utils.GetCache("gzh_access_token")
	if _token_info == nil {
		_rs, _err := g.Client().Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", this.AppID, this.Secret))
		if _err == nil {
			_get_json := _rs.ReadAllString()
			glog.Println("直接接口获取token", _get_json)
			utils.SetCache("gzh_access_token", gconv.Map(_get_json), 7000)
		}
		_token_info, _ = utils.GetCache("gzh_access_token")
	} else {
		glog.Println("从缓存获取token", gconv.Map(_token_info))
	}
	_accessToken.AccessToken = gconv.String(gconv.Map(_token_info)["access_token"])
	_accessToken.ExpiresIN = gconv.Duration(gconv.Map(_token_info)["expires_in"])
	return _accessToken
}

/**
 * @Description: 获取微信API接口 IP地址
 * @receiver this
 */
func (this *GzhApi) GetApiDomainIP() (interface{}, error) {
	var _rt interface{}
Label1:
	_token_info := this.get_token()
	_rs, _err := g.Client().Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/get_api_domain_ip?access_token=%s", _token_info.AccessToken))
	if _err == nil {
		_get_json := _rs.ReadAllString()
		if strings.Contains(_get_json, "40001") {
			utils.RemoveCache("gzh_access_token")
			goto Label1
		} else {
			_rt = gconv.Map(_get_json)
		}
	}
	return _rt, nil
}

/**
 * @Description: 获取微信callback IP地址
 * @receiver this
 */
func (this *GzhApi) GetCallBackIP() (interface{}, error) {
	var _rt interface{}
Label1:
	_token_info := this.get_token()
	_rs, _err := g.Client().Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/getcallbackip?access_token=%s", _token_info.AccessToken))
	if _err == nil {
		_get_json := _rs.ReadAllString()
		if strings.Contains(_get_json, "40001") {
			utils.RemoveCache("gzh_access_token")
			goto Label1
		} else {
			_rt = gconv.Map(_get_json)
		}
	}
	return _rt, nil
}

/**
 * @Description: 通过code换取网页授权access_token
 * @receiver this
 */
func (this *GzhApi) CodeAccessToken(_code, _grant_type string) (interface{}, error) {
	var _rt interface{}
	if _grant_type == "" {
		_grant_type = "authorization_code"
	}
	_rs, _err := g.Client().Get(fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=%s", this.AppID, this.Secret, _code, _grant_type))
	if _err == nil {
		_rt = gconv.Map(_rs.ReadAllString())
	}
	return _rt, nil
}
