package tt_minigame_api

import (
	"fmt"
	"resty.dev/v3"
	"time"
)

type AppsV2TokenRequest struct {
	GrantType *string `json:"grant_type,omitempty" xml:"grant_type,omitempty" require:"true"`
	Appid     *string `json:"appid,omitempty" xml:"appid,omitempty" require:"true"`
	Secret    *string `json:"secret,omitempty" xml:"secret,omitempty" require:"true"`
}
type AppsV2TokenResponseData struct {
	AccessToken *string `json:"access_token,omitempty" xml:"access_token,omitempty" require:"true"`
	ExpiresIn   *int64  `json:"expires_in,omitempty" xml:"expires_in,omitempty" require:"true"`
	ExpiresAt   *int64  `json:"expires_at,omitempty" xml:"expires_at,omitempty" require:"true"`
}

func (s *AppsV2TokenRequest) SetGrantType(v string) *AppsV2TokenRequest {
	s.GrantType = &v
	return s
}

func (s *AppsV2TokenRequest) SetAppid(v string) *AppsV2TokenRequest {
	s.Appid = &v
	return s
}

func (s *AppsV2TokenRequest) SetSecret(v string) *AppsV2TokenRequest {
	s.Secret = &v
	return s
}

type AppsV2TokenResponse struct {
	ErrTips string                   `json:"err_tips,omitempty" xml:"err_tips,omitempty" require:"true"`
	Data    *AppsV2TokenResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo   int64                    `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s *AppsV2TokenResponseData) SetAccessToken(v string) *AppsV2TokenResponseData {
	s.AccessToken = &v
	return s
}

func (s *AppsV2TokenResponseData) SetExpiresIn(v int64) *AppsV2TokenResponseData {
	s.ExpiresIn = &v
	return s
}

func (s *AppsV2TokenResponseData) SetExpiresAt(v int64) *AppsV2TokenResponseData {
	s.ExpiresAt = &v
	return s
}

func (t *TT) GetAccessToken(request *AppsV2TokenRequest, url string) (_result *AppsV2TokenResponse, _err error) {
	if url == "" {
		url = "https://minigame.zijieapi.com/mgplatform/api/apps/v2/token"
	}
	client := resty.New()
	defer client.Close()
	resp, err := client.R().SetHeader("Content-Type", "application/json").
		SetBody(request).SetRetryCount(t.MaxAttempts).SetRetryWaitTime(1 * time.Second).
		SetResult(_result).
		Post(url)
	if err != nil {
		return nil, fmt.Errorf("get getAccessToken failed, unknow err, err=%w", err)
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("get getClientToken failed, http statusCode not 200, code=%d", resp.StatusCode())
	}

	if _result.ErrNo != 0 {
		return nil, fmt.Errorf("get getClientToken failed, biz code not 0, err_no=%d err_msg=%s", _result.ErrNo, _result.ErrTips)
	}
	_err = err
	tt.ClientToken = &CommonToken{
		ExpireIn:  *_result.Data.ExpiresIn,
		StartTime: time.Now().Add(-5 * time.Second),
		Token:     *_result.Data.AccessToken,
	}
	return
}

type AppsCode2sessionRequest struct {
	Appid         *string `json:"appid,omitempty" xml:"appid,omitempty" require:"true"`
	Secret        *string `json:"secret,omitempty" xml:"secret,omitempty" require:"true"`
	Code          *string `json:"code,omitempty" xml:"code,omitempty"`
	AnonymousCode *string `json:"anonymous_code,omitempty" xml:"anonymous_code,omitempty"`
}

func (s *AppsCode2sessionRequest) SetAppid(v string) *AppsCode2sessionRequest {
	s.Appid = &v
	return s
}

func (s *AppsCode2sessionRequest) SetSecret(v string) *AppsCode2sessionRequest {
	s.Secret = &v
	return s
}

func (s *AppsCode2sessionRequest) SetCode(v string) *AppsCode2sessionRequest {
	s.Code = &v
	return s
}

func (s *AppsCode2sessionRequest) SetAnonymousCode(v string) *AppsCode2sessionRequest {
	s.AnonymousCode = &v
	return s
}

type AppsCode2sessionResponse struct {
	Error           int64   `json:"error,omitempty" xml:"error,omitempty" require:"true"`
	Errcode         int64   `json:"errcode,omitempty" xml:"errcode,omitempty"`
	SessionKey      *string `json:"session_key,omitempty" xml:"session_key,omitempty"`
	Errmsg          string  `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	Openid          *string `json:"openid,omitempty" xml:"openid,omitempty"`
	Message         *string `json:"message,omitempty" xml:"message,omitempty"`
	Unionid         *string `json:"unionid,omitempty" xml:"unionid,omitempty"`
	AnonymousOpenid *string `json:"anonymous_openid,omitempty" xml:"anonymous_openid,omitempty"`
}

func (s *AppsCode2sessionResponse) SetSessionKey(v string) *AppsCode2sessionResponse {
	s.SessionKey = &v
	return s
}

func (s *AppsCode2sessionResponse) SetOpenid(v string) *AppsCode2sessionResponse {
	s.Openid = &v
	return s
}

func (s *AppsCode2sessionResponse) SetMessage(v string) *AppsCode2sessionResponse {
	s.Message = &v
	return s
}

func (s *AppsCode2sessionResponse) SetUnionid(v string) *AppsCode2sessionResponse {
	s.Unionid = &v
	return s
}

func (s *AppsCode2sessionResponse) SetAnonymousOpenid(v string) *AppsCode2sessionResponse {
	s.AnonymousOpenid = &v
	return s
}
func (t *TT) Code2Session(request *AppsCode2sessionRequest, url string) (_result *AppsCode2sessionResponse, _err error) {
	if url == "" {
		url = "https://minigame.zijieapi.com/mgplatform/api/apps/jscode2session"
	}
	client := resty.New()
	defer client.Close()
	resp, err := client.R().SetHeader("Content-Type", "application/json").
		SetQueryParams(map[string]string{
			"appid":          *request.Appid,
			"secret":         *request.Secret,
			"code":           *request.Code,
			"anonymous_code": *request.AnonymousCode,
		}).SetRetryCount(t.MaxAttempts).SetRetryWaitTime(1 * time.Second).
		SetResult(_result).
		Get(url)
	if err != nil {
		return nil, fmt.Errorf("get Code2Session failed, unknow err, err=%w", err)
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("get Code2Session failed, http statusCode not 200, code=%d", resp.StatusCode())
	}

	if _result.Error != 0 {
		return nil, fmt.Errorf("get Code2Session failed, biz code not 0,error=%d, err_no=%d err_msg=%s", _result.Error, _result.Errcode, _result.Errmsg)
	}
	_err = err
	return
}
