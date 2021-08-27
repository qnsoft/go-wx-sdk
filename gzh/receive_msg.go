package gzh

/*---------------------------------------------------普通消息--------------------------------------------------------*/
/**
 * @Description: 消息对象基础类
 */
type MsgBase struct {
	//开发者微信号
	ToUserName string `xml:"ToUserName" json:"ToUserName"`
	//发送方帐号（一个OpenID）
	FromUserName string `xml:"FromUserName" json:"FromUserName"`
	//消息创建时间 （整型）
	CreateTime string `xml:"CreateTime" json:"CreateTime"`
	//消息类型，文本为text,消息类型，图片为image,语音为voice,视频为video,小视频为shortvideo,地理位置为location,链接为link，事件event,扫描 SCAN，上报地理位置 LOCATION
	MsgType string `xml:"MsgType" json:"MsgType"`
}

/**
 * @Description: 消息对象普通消息基础类
 */
type MsgPT struct {
	MsgBase
	//消息id，64位整型
	MsgID string `xml:"MsgId" json:"MsgId"`
}

/**
 * @Description: 消息对象-文本消息
 */
type MsgText struct {
	MsgPT
	//文本消息内容
	Content string `xml:"Content" json:"Content"`
}

/**
 * @Description: 消息对象-图片消息
 */
type MsgPic struct {
	MsgPT
	//
	PicUrl string `xml:"PicUrl" json:"PicUrl"`
	//图片消息媒体id，可以调用获取临时素材接口拉取数据。
	MediaId string `xml:"MediaId" json:"MediaId"`
}

/**
 * @Description: 消息对象-语音消息
 */
type MsgMedia struct {
	MsgPT
	//语音消息媒体id，可以调用获取临时素材接口拉取数据。
	MediaId string `xml:"MediaId" json:"MediaId"`
	//语音格式，如amr，speex等
	Format string `xml:"Format" json:"Format"`
	//语音识别结果，UTF8编码
	Recognition string `xml:"Recognition" json:"Recognition"`
}

/**
 * @Description: 消息对象-视频消息
 */
type MsgVideo struct {
	MsgPT
	//视频消息媒体id，可以调用获取临时素材接口拉取数据。
	MediaId string `xml:"MediaId" json:"MediaId"`
	//视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。
	ThumbMediaId string `xml:"ThumbMediaId" json:"ThumbMediaId"`
}

/**
 * @Description: 消息对象-小视频消息
 */
type MsgShortVideo struct {
	MsgPT
	//视频消息媒体id，可以调用获取临时素材接口拉取数据。
	MediaId string `xml:"MediaId" json:"MediaId"`
	//视频消息缩略图的媒体id，可以调用获取临时素材接口拉取数据。
	ThumbMediaId string `xml:"ThumbMediaId" json:"ThumbMediaId"`
}

/**
 * @Description: 消息对象-地理位置消息
 */
type MsgLocation struct {
	MsgPT
	//地理位置纬度
	LocationX string `xml:"Location_X" json:"Location_X"`
	//地理位置经度
	LocationY string `xml:"Location_Y" json:"Location_Y"`
	//地图缩放大小
	Scale string `xml:"Scale" json:"Scale"`
	//地理位置信息
	Label string `xml:"Label" json:"Label"`
}

/**
 * @Description: 消息对象-链接消息
 */
type MsgLink struct {
	MsgPT
	//消息标题
	Title string `xml:"Title" json:"Title"`
	//消息描述
	Description string `xml:"Description" json:"Description"`
	//消息链接
	Url string `xml:"Url" json:"Url"`
}

/*---------------------------------------------------事件消息--------------------------------------------------------*/
/**
 * @Description:取消关注事件
 */
type MsgSubscribeOrUnsubscribe struct {
	MsgBase
	//事件类型，subscribe(订阅)、unsubscribe(取消订阅)
	Event string `xml:"Event" json:"Event"`
}

/**
 * @Description:扫描带参数二维码事件
 */
type MsgScan struct {
	MsgBase
	//事件类型，subscribe
	Event string `xml:"Event" json:"Event"`
	//事件KEY值，qrscene_为前缀，后面为二维码的参数值
	EventKey string `xml:"EventKey" json:"EventKey"`
	//二维码的ticket，可用来换取二维码图片
	Ticket string `xml:"Ticket" json:"Ticket"`
}

/**
 * @Description:用户已关注时的事件推送
 */
type MsgSubscribed struct {
	MsgBase
	//事件类型，SCAN
	Event string `xml:"Event" json:"Event"`
	//事件KEY值，是一个32位无符号整数，即创建二维码时的二维码scene_id
	EventKey string `xml:"EventKey" json:"EventKey"`
	//二维码的ticket，可用来换取二维码图片
	Ticket string `xml:"Ticket" json:"Ticket"`
}

/**
 * @Description:上报地理位置信息
 */
type MsgPostLocation struct {
	MsgBase
	//事件类型，LOCATION
	Event string `xml:"Event" json:"Event"`
	//地理位置纬度
	Latitude string `xml:"Latitude" json:"Latitude"`
	//地理位置经度
	Longitude string `xml:"Longitude" json:"Longitude"`
	//地理位置精度
	Precision string `xml:"Precision" json:"Precision"`
}

/**
 * @Description:自定义菜单事件
 */
type MsgCustomMenu struct {
	MsgBase
	//事件类型，CLICK
	Event string `xml:"Event" json:"Event"`
	//事件KEY值，与自定义菜单接口中KEY值对应
	EventKey string `xml:"EventKey" json:"EventKey"`
}

/*---------------------------------------------------被动回复用户消息--------------------------------------------------------*/
/**
 * @Description:回复文本消息
 */
type ReplyMsgText struct {
	MsgBase
	//回复的消息内容（换行：在content中能够换行，微信客户端就支持换行显示）
	Content string `xml:"Content" json:"Content"`
}

/**
 * @Description:回复图片消息
 */
type ReplyMsgPic struct {
	MsgBase
	//回复的图片消息内容
	Image struct {
		//通过素材管理中的接口上传多媒体文件，得到的id
		MediaId string `xml:"MediaId" json:"MediaId"`
	} `xml:"Image" json:"Image"`
}

/**
 * @Description:回复语音消息
 */
type ReplyMsgMedia struct {
	MsgBase
	//回复的语音消息内容
	Voice struct {
		//通过素材管理中的接口上传多媒体文件，得到的id
		MediaId string `xml:"MediaId" json:"MediaId"`
	} `xml:"Voice" json:"Voice"`
}

/**
 * @Description:回复视频消息
 */
type ReplyMsgVideo struct {
	MsgBase
	//回复的视频消息内容
	Video struct {
		//通过素材管理中的接口上传多媒体文件，得到的id
		MediaId string `xml:"MediaId" json:"MediaId"`
		//视频消息的标题
		Title string `xml:"Title" json:"Title"`
		//视频消息的描述
		Description string `xml:"MediaId" json:"Description"`
	} `xml:"Video" json:"Video"`
}

/**
 * @Description:回复音乐消息
 */
type ReplyMsgMusic struct {
	MsgBase
	//回复的音乐消息内容
	Music struct {
		//音乐标题
		Title string `xml:"Title" json:"Title"`
		//音乐描述
		Description string `xml:"MediaId" json:"Description"`
		//音乐链接
		MusicUrl string `xml:"MusicUrl" json:"MusicUrl"`
		//高质量音乐链接，WIFI环境优先使用该链接播放音乐
		HQMusicUrl string `xml:"MediaId" json:"HQMusicUrl"`
		//缩略图的媒体id，通过素材管理中的接口上传多媒体文件，得到的id
		ThumbMediaId string `xml:"MediaId" json:"ThumbMediaId"`
	} `xml:"Music" json:"Music"`
}

/**
 * @Description:回复图文消息
 */
type ReplyMsgArticles struct {
	MsgBase
	//图文消息个数；当用户发送文本、图片、语音、视频、图文、地理位置这六种消息时，开发者只能回复1条图文消息；其余场景最多可回复8条图文消息
	ArticleCount string `xml:"ArticleCount" json:"ArticleCount"`
	//图文消息信息，注意，如果图文数超过限制，则将只发限制内的条数
	Articles []struct {
		Item struct {
			//图文消息标题
			Title string `xml:"Title" json:"Title"`
			//图文消息描述
			Description string `xml:"MediaId" json:"Description"`
			//图片链接，支持JPG、PNG格式，较好的效果为大图360*200，小图200*200
			PicUrl string `xml:"PicUrl" json:"PicUrl"`
			//点击图文消息跳转链接
			Url string `xml:"Url" json:"Url"`
		} `xml:"item" json:"item"`
	} `xml:"Articles" json:"Articles"`
}

/*---------------------------------------接收消息----------------------------------------------------------*/
/**
 * @Description:设置所属行业(小心使用，一旦使用该接口线上模板会被自动清空)
 * @receiver this
 */
//func (this *GzhApi) SetIndustry() (interface{}, error) {
//	var _rt interface{}
//	_token_info := this.get_token()
//	_rs, _err := g.Client().Post(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/template/api_set_industry?access_token=%s", _token_info.AccessToken), gconv.String(g.Map{
//		"industry_id1": "2",
//		"industry_id2": "4",
//	}))
//	if _err == nil {
//		_rt = gconv.Map(_rs.ReadAllString())
//	}
//	return _rt, nil
//}
