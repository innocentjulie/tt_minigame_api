# tt_minigame_api
tt_minigame_api some api 抖音 小游戏部分api

## 安装
    go get github.com/innocentjulie/tt_minigame_api

## 使用
```go
	tt := tt_minigame_api.GetTT()
	grandType := "client_credential"
	appId := "tt4233**"
	secret := "cbs***"
	req := &tt_minigame_api.AppsV2TokenRequest{
		GrantType: &grandType,
		Appid:     &appId,
		Secret:    &secret,
	}
	result, err := tt.GetAccessToken(req, "")
	if err != nil {
		fmt.Errorf(fmt.Sprintln("get access token err:", err))
		return
	}
	fmt.Println("get access token success")
	fmt.Println(result)
```

## 更新
 - v0.0.1 新增getAccessToken和code2session
