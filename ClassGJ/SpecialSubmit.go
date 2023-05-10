package ClassGJ

import (
	"fmt"
	"github.com/imroc/req/v3"
)

func SpecialSubmit(title string, workid string, filepath string, member_id string, imprint string) {
	requestbody := map[string]interface{}{
		"extra":       1,
		"id":          workid,
		"daka_day":    "",
		"submit_type": "submit",
		"networkType": "wifi",
		"member_id":   member_id,
		"op":          "add",
		"files": []map[string]interface{}{
			{
				"file":     "wxfile://temp/test.png",
				"cate":     "img",
				"new_name": filepath,
				"type":     1,
				"size":     525128,
				"uploaded": true,
			},
		},
		"feedback_text": "",
	}
	var response Code
	client := req.C()
	_, _ = client.R().
		SetBodyJsonMarshal(requestbody).
		SetHeaders(map[string]string{
			"Imprint":         imprint,
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 MicroMessenger/7.0.20.1781(0x6700143B) NetType/WIFI MiniProgramEnv/Windows WindowsWechat/WMPF XWEB/6500",
			"Content-Type":    "application/json",
			"Accept-Language": "zh-CN,zh",
		}).
		SetSuccessResult(&response).
		Post("https://a.welife001.com/applet/notify/feedbackWithOss")
	if response.Code == 0 {
		fmt.Println("成功完成" + title + "青年大学习!")
	}
}
