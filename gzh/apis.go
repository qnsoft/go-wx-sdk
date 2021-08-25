package gzh

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/qnsoft/go-wx-sdk/utils"
)

var (
	//appid
	AppID string = "wx24979f86c90e4e65"
	//secret
	Secret string = "695dfbd7e0f83120d3f45b07a3a10c82"
)

/*
公众号对象
*/
type GzhApi struct {
}

/**
 * @Description: 获取token
 * @return *AccessTokenModel
 */
func (this *GzhApi) get_token() AccessTokenModel {
	_accessToken := AccessTokenModel{}
	_token_model, _token_err := utils.GetCache("access_token", "", 700)
	if _token_err == nil && _token_model != nil {
		_errA := gconv.Struct(_token_model, &_accessToken)
		if _errA != nil {
			_rs, _err := g.Client().Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", AppID, Secret))
			if _err == nil {
				_err1 := gconv.Struct(_rs.ReadAllString(), &_accessToken)
				if _err1 == nil {
					utils.SetCache("access_token", _accessToken, _accessToken.ExpiresIN)
				}
			}
		}
	} else {
		_rs, _err := g.Client().Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", AppID, Secret))
		if _err == nil {
			_err1 := gconv.Struct(_rs.ReadAllString(), &_accessToken)
			if _err1 == nil {
				utils.SetCache("access_token", _accessToken, _accessToken.ExpiresIN)
			}
		}
	}
	return _accessToken
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
	_rs, _err := g.Client().Get(fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=%s", AppID, Secret, _code, _grant_type))
	if _err == nil {
		_rt = gconv.Map(_rs.ReadAllString())
	}
	return _rt, nil
}

//---------------------------------------模板消息----------------------------------------------------------
/**
 * @Description:设置所属行业(小心使用，一旦使用该接口线上模板会被自动清空)
 * @receiver this
 */
func (this *GzhApi) SetIndustry() (interface{}, error) {
	var _rt interface{}
	_token_info := this.get_token()
	_rs, _err := g.Client().Post(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/template/api_set_industry?access_token=%s", _token_info.AccessToken), gconv.String(g.Map{
		"industry_id1": "2",
		"industry_id2": "4",
	}))
	if _err == nil {
		_rt = gconv.Map(_rs.ReadAllString())
	}
	return _rt, nil
}

/**
 * @Description:获取设置的行业信息
 * @receiver this
 */
func (this *GzhApi) GetIndustry() (interface{}, error) {
	var _rt interface{}
	_token_info := this.get_token()
	_rs, _err := g.Client().Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/template/get_industry?access_token=%s", _token_info.AccessToken))
	if _err == nil {
		_rt = gconv.Map(_rs.ReadAllString())
	}
	return _rt, _err
}

/**
 * @Description: 获得模板ID
 * @receiver this
 * @param _template_id_short 模板库中模板的编号，有“TM**”和“OPENTMTM**”等形式
 * @return interface{}
 * @return error
 */
func (this *GzhApi) AddTemplate(_template_id_short string) (interface{}, error) {
	var _rt interface{}
	_token_info := this.get_token()
	_rs, _err := g.Client().Post(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/template/api_add_template?access_token=%s", _token_info.AccessToken), gconv.String(g.Map{
		"template_id_short": _template_id_short,
	}))
	if _err == nil {
		_rt = gconv.Map(_rs.ReadAllString())
	}
	return _rt, _err
}

/**
 * @Description: 获取模板列表
 * @receiver this
 * @param _template_id_short
 * @return interface{}
 * @return error
 */
func (this *GzhApi) GetAllPrivateTemplate() (interface{}, error) {
	var _rt interface{}
	_token_info := this.get_token()
	_rs, _err := g.Client().Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/template/get_all_private_template?access_token=%s", _token_info.AccessToken))
	if _err == nil {
		_rt = gconv.Map(_rs.ReadAllString())
	}
	return _rt, _err
}

/**
 * @Description:删除模板
 * @receiver this
 * @param _template_id 模板id
 * @return interface{}
 * @return error
 */
func (this *GzhApi) DelPrivateTemplate(_template_id string) (interface{}, error) {
	var _rt interface{}
	_token_info := this.get_token()
	_rs, _err := g.Client().Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/template/del_private_template?access_token=%s", _token_info.AccessToken), gconv.String(g.Map{
		"template_id": _template_id,
	}))
	if _err == nil {
		_rt = gconv.Map(_rs.ReadAllString())
	}
	return _rt, _err
}

/**
 * @Description:发送模板消息
 * @receiver this
 * @param _map 模板消息内容
 * @return interface{}
 * @return error
 */
func (this *GzhApi) SendTemplate(_map g.Map) (interface{}, error) {
	var _rt interface{}
	_token_info := this.get_token()
	_rs, _err := g.Client().Post(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", _token_info.AccessToken), gconv.String(_map))
	if _err == nil {
		_rt = gconv.Map(_rs.ReadAllString())
	}
	return _rt, _err
}
