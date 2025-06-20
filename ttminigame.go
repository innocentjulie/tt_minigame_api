package tt_minigame_api

import "time"

type TT struct {
	AutoRetry      bool
	IgnoreSSL      bool
	MaxAttempts    int
	ReadTimeout    int
	ConnectTimeout int
	ClientToken    *CommonToken //保存用户获得的token token 是小游戏级别 token，不要为每个用户单独分配一个 token，会导致 token 校验失败。建议每小时更新一次即可。
}

type CommonToken struct {
	Token     string
	StartTime time.Time
	ExpireIn  int64
}

// 将TT 弄成单例返回,避免多次调用多个token被互刷的情况
var tt *TT

func GetTT() *TT {
	if tt != nil {
		return tt
	}
	tt = GetDefaultTT()
	return tt
}

// RegisterTT 注册自定义配置的TT
func RegisterTT(t *TT) {
	tt = t
}

func GetDefaultTT() *TT {
	return &TT{
		AutoRetry:      true,
		IgnoreSSL:      false,
		MaxAttempts:    5,
		ReadTimeout:    5000,
		ConnectTimeout: 1000,
	}
}

type Config struct {
	ClientKey    *string `json:"clientKey,omitempty" xml:"clientKey,omitempty" require:"true"`
	ClientSecret *string `json:"clientSecret,omitempty" xml:"clientSecret,omitempty" require:"true"`
}

func NewConfig(key string, secret string) *Config {
	return &Config{
		ClientKey:    &key,
		ClientSecret: &secret,
	}
}

func (s *Config) SetClientKey(v string) *Config {
	s.ClientKey = &v
	return s
}

func (s *Config) SetClientSecret(v string) *Config {
	s.ClientSecret = &v
	return s
}

type Token struct {
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty" require:"true"`
	ExpiresIn   *int64  `json:"expiresIn,omitempty" xml:"expiresIn,omitempty" require:"true"`
}

func (s *Token) SetAccessToken(v string) *Token {
	s.AccessToken = &v
	return s
}

func (s *Token) SetExpiresIn(v int64) *Token {
	s.ExpiresIn = &v
	return s
}
