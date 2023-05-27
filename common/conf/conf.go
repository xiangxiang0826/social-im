package conf

type RateLimitConfig struct {
	Seconds int
	Quota   int
}

type JwtAuth struct {
	AccessSecret string
	AccessExpire int64
}

type AgoraConf struct {
	AppId          string
	AppCertificate string
	Server         string
	BackupServer   string
	PubChannel     string
}

type AliOssConf struct {
	AccessKeyId      string
	AccessKeySecrect string
	SignName         string
	VerifyTemplate   string
	Enable           bool
}
