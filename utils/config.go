package utils

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

/*
  default_module: 'api'
  weixin:
    appid: '' #小程序 appid
    secret: '' #小程序密钥
    mch_id: '' #商户帐号ID
    partner_key: '' #微信支付密钥
    notify_url: '' #微信异步通知，例：https://www.nideshop.com/api/pay/notify
  express:
    #快递物流信息查询使用的是快递鸟接口，申请地址：http://www.kdniao.com/
    appid: ''  #对应快递鸟用户后台 用户ID
    appkey: '' #对应快递鸟用户后台 API key
    request_url: 'http://api.kdniao.cc/Ebusiness/EbusinessOrderHandle.aspx'
*/

type Config struct {
	Default_moudle string  `yaml:"default_module"`
	Wx             WeiXin  `yaml:"weixin"`
	Exp            Express `yaml:"express"`
}

type WeiXin struct {
	Appid       string `yaml:"appid"`
	Secret      string `yaml:"secret"`
	Mch_id      string `yaml:"mch_id"`
	Partner_key string `yaml:"partner_key"`
	Notify_url  string `yaml:"notify_url"`
}
type Express struct {
	Appid       string `yaml:"appid"`
	Appkey      string `yaml:"appkey"`
	Request_url string `yaml:"request_url"`
}

var configFile []byte

func init() {
	var err error
	configFile, err = ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("yamlFile.Get err %v ", err)
	}
}

func GetConfig() (e *Config, err error) {
	err = yaml.Unmarshal(configFile, &e)
	return e, err
}
