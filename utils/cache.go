package utils

import (
	"github.com/gogf/gf/os/gcache"
	"time"
)
/**
 * @Description: 设置缓存
 * @param _key 缓存键
 * @param _value 缓存值
 * @param expires_in 有效时间，单位：秒
 */
func SetCache(_key,_value interface{},expires_in time.Duration)(bool,error){
	// 当键名不存在时写入，设置过期时间1000毫秒
	_bl,_err :=gcache.SetIfNotExist(_key, _value, expires_in*1000*time.Millisecond)
	return _bl,_err
}
/**
 * @Description: 获取缓存
 * @param _key 缓存键
 * @param _value 缓存值
 * @param expires_in 有效时间，单位：秒
 */
func GetCache(_key,_value interface{},expires_in time.Duration)(interface{},error){
	res,err:= gcache.GetOrSet(_key,_value,expires_in)
	return res,err

}
/**
 * @Description: 根据键删除缓存
 * @param _key
 * @return interface{}
 * @return error
 */
func RemoveCache(_key interface{})(interface{},error){
	res,err:=gcache.Remove(_key)
	return res,err
}
