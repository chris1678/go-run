package config

type Alioss struct {
	AccessKeyId     string // 请填写您的AccessKeyId。
	AccessKeySecret string // 请填写您的AccessKeySecret。
	Endpoint        string // host的格式为 bucketname.endpoint ，请替换为您的真实信息。
	CallbackUrl     string // callbackUrl为 上传回调服务器的URL，请将下面的IP和Port配置为您自己的真实信息。
	UploadDir       string // 用户上传文件时指定的前缀。
	ExpireTime      int64
}

var AliossConfig = new(Alioss)
