package tests

import (
	"github.com/zhaoyunxing92/dingtalk/domain"
	"testing"
)

//发送工作通知
func TestSendTextWorkNotify(t *testing.T) {
	res := domain.NewTextWorkNotify("hello dingtalk")
	res.UserIds = []string{"manager164"}
	notify, err := dingTalk.SendWorkNotify(res)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(notify)
}

func TestSendOAWorkNotify(t *testing.T) {

	oa := domain.OA{}

	f := domain.Form{}
	f.Key = "姓名:"
	f.Value = "赵云兴"

	f2 := domain.Form{}
	f2.Key = "年龄:"
	f2.Value = "20"

	f3 := domain.Form{}
	f3.Key = "身高:"
	f3.Value = "190"

	f4 := domain.Form{}
	f4.Key = "体重:"
	f4.Value = "80KG"

	f5 := domain.Form{}
	f5.Key = "爱好:"
	f5.Value = "go、java、docker、vue"

	//设置体
	oa.Body.Forms = append(oa.Body.Forms, f, f2, f3, f4, f5)
	oa.Body.Content = "消息体的标题，建议50个字符以内。"
	oa.Body.Title = "头部标题"
	oa.Body.Author = "赵云兴"

	//设置头
	oa.Header.BgColor = "FFBBBBBB"
	oa.Header.Text = "被替换为应用名称"

	oa.MessageUrl = "https://ding-doc.dingtalk.com/document#/org-dev-guide/message-types-and-data-format"
	oa.PcMessageUrl = "https://ding-doc.dingtalk.com"


	res := domain.NewOAWorkNotify(oa)
	res.UserIds = []string{"manager164"}

	notify, err := dingTalk.SendWorkNotify(res)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v",notify)
}