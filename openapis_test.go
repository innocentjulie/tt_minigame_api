package tt_minigame_api

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDefaultTT(t *testing.T) {
	type args struct { //用到的参数值
		userId   int
		showType int
	}
	tests := []struct {
		name    string                    //测试用例名称
		args    args                      //测试参数
		want    *TT                       //期望结果
		wantErr assert.ErrorAssertionFunc //错误断言
	}{
		{
			name: "test1",
			args: args{userId: 5, showType: 6},
			want: GetTT(),
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Nil(t, err, i)
			},
		},
		//{
		//	name: "test2",
		//	args: args{userId: 5, showType: 5},
		//	want: model.UserShowRoomResp{},
		//	wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
		//		return assert.Nil(t, err, i)
		//	},
		//},
		//{
		//	name: "test3",
		//	args: args{userId: 5, showType: 3},
		//	want: model.UserShowRoomResp{},
		//	wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
		//		return assert.Nil(t, err, i)
		//	},
		//},
	}
	for _, tmp := range tests {
		t.Run(tmp.name, func(t *testing.T) {
			got := GetDefaultTT()
			//if !tt.wantErr(t, err, fmt.Sprintf("GetShowroomListByType() error should be equal,but got %v", err)) {
			//	return
			//}
			if got != nil {
				assert.Equal(t, tmp.want, got, "tt obj should be equal")
				assert.Equal(t, 5, got.MaxAttempts, "default maxAttempts  should be 5")
				assert.Equal(t, 1000, got.ConnectTimeout, "default connectTimeout should be 1000")
			}
		})
	}
}

func TestTT_GetAccessToken(t *testing.T) {
	tt := GetTT()
	grandType := "client_credential"
	appId := ""
	secret := ""
	req := &AppsV2TokenRequest{
		GrantType: &grandType,
		Appid:     &appId,
		Secret:    &secret,
	}
	result, err := tt.GetAccessToken(req, "")
	if err != nil {
		fmt.Println(err)
		fmt.Sprintln("get access token err:", err)
		return
	}
	fmt.Println("get access token success")
	fmt.Println(result)
}
func TestTT_Code2Session(t *testing.T) {

}
